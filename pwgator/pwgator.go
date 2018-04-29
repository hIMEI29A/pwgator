// Copyright 2018 hIMEI
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pwgator

import (
	"strings"
)

// Leet is a map for string literals and its leet-coded aliases
var Leet = map[string][]string{
	"A": {"a", "@", "*", "A"},
	"E": {"e", "3", "E"},
	"S": {"s", "$", "5", "S", "2"},
	"O": {"o", "0", "O"},
	"T": {"7", "t", "T"},
	"F": {"f", "F", "4"},
	"I": {"i", "!", "1", "I"},
	"B": {"b", "8", "6", "B"},
	"G": {"g", "9", "G"},
}

// Leet "leetcodes" given string
func leet(pass string) string {
	var newpass string

	for i := range pass {
		if _, ok := Leet[string(pass[i])]; ok {
			values := Leet[string(pass[i])]
			trand := ranint(len(values))
			newpass = newpass + values[trand]
		} else {
			newpass = newpass + string(pass[i])
		}
	}

	return newpass
}

// Token is a type for string literal's representation
type Token struct {
	// Token is a field for integer value which may by converted to enum
	Token int
	// TokenType is a field for enum type
	TokenType TOKEN_T
}

// NewToken instantiates Token type
func NewToken() *Token {
	token := &Token{}

	token.TokenType = TOKEN_T(ranint(len(Tokens)))
	token.Token = ranint(len(token.TokenType.TokenNames()))

	return token
}

// String gives string representation of Token
func (t *Token) String() string {
	var tokenized string

	switch t.TokenType {
	case CONSONANT:
		tokenized = CONSONANT_T(t.Token).String()

	case VOWEL:
		tokenized = VOWEL_T(t.Token).String()

	case DIPTHONG:
		tokenized = DIPTHONG_T(t.Token).String()

	}

	return tokenized
}

// Secret is a type for set of string literals wich will be generated
type Secret struct {
	// Generated tokens
	Tokens []*Token
	// Desired length
	Length int
}

//NewSecret instantiates Secret type
func NewSecret() *Secret {
	secret := &Secret{}

	return secret
}

func (s *Secret) nextToken(num int) *Token {
	t := &Token{}

	if num < len(s.Tokens)-1 {
		t = s.Tokens[num+1]
	} else {
		t = s.Tokens[len(s.Tokens)-1]
	}
	return t
}

func (s *Secret) insertToken(token *Token, num int) []*Token {
	newtoken := &Token{}
	newtokens := s.Tokens

	newtokens = append(newtokens, newtoken)
	copy(newtokens[num+1:], newtokens[num:])
	newtokens[num] = token

	return newtokens
}

func (s *Secret) generateTokens(length int) {
	var secret []*Token

	for i := 0; i < length; i++ {
		token := NewToken()
		secret = append(secret, token)
	}
	s.Length = len(secret)

	s.Tokens = secret
}

// typeToken checks if given token is an object
// of given token type (ENUM). It useful for checking DIPHTHONG for vowels
func (s *Secret) typeToken(ttype TOKEN_T, token *Token) bool {
	check := false

	if token.TokenType == ttype {
		check = true
	}

	if token.TokenType == DIPTHONG {
		tstring := token.String()

		if ttype == CONSONANT {
			if Consonant_t(string(tstring[0])) == true && Consonant_t(string(tstring[1])) == true {
				check = true
			}
		}

		if ttype == VOWEL {
			if Vowel_t(string(tstring[0])) == true && Vowel_t(string(tstring[1])) == true {
				check = true
			}
		}
	}

	return check
}

// tune makes string manipulations with generated data
func (s *Secret) tune(val string) string {
	word := downize(leet(val))
	var newword string

	if len(word) > s.Length {
		for i := 0; i < s.Length; i++ {
			newword = newword + string(word[i])
		}
	} else {
		newword = word
	}

	return newword
}

// parse makes manipulation with Secret's []*Token
func (s *Secret) parse() string {
	newtokens := s.Tokens

	for i := range newtokens {

		if s.typeToken(VOWEL, newtokens[i]) == true && s.typeToken(VOWEL, s.nextToken(i)) == true {
			newtoken := NewToken()
			newtoken.Token = int(consonant_t())
			newtoken.TokenType = CONSONANT

			newtokens = s.insertToken(newtoken, i+1)
		}

		if s.typeToken(CONSONANT, newtokens[i]) == true && s.typeToken(CONSONANT, s.nextToken(i)) == true {
			newtoken := NewToken()
			newtoken.Token = int(vowel_t())
			newtoken.TokenType = VOWEL

			newtokens = s.insertToken(newtoken, i+1)
		}
	}

	s.Tokens = newtokens

	return s.String()
}

// PassWord generates password with given length and
// given mode (strong or humanized)
func (s *Secret) PassWord(length int, strong bool) string {
	var p string

	if !strong {
		s.generateTokens(length)
		p = s.parse()
	} else {
		p = genstrong(length)
	}

	return p
}

// PassPhrase generates passphrase with given length and
// given mode (strong or humanized)
func (s *Secret) PassPhrase(words int, strong, random bool) string {
	var (
		phrase   string
		template []int
	)

	template = GetPhraseTemplate(words)

	if random == true {
		template = GetRandomTemplate(words)
	}

	if !strong {
		for i := range template {
			phrase = phrase + SPACE + s.PassWord(template[i], strong)
		}
	} else {
		for i := range template {
			phrase = phrase + SPACE + genstrong(template[i])
		}
	}

	return strings.TrimPrefix(phrase, SPACE)
}

// String stringizes Secret's []*Token
func (s *Secret) String() string {
	var sec string

	for i := range s.Tokens {
		sec = sec + s.Tokens[i].String()
	}

	return s.tune(sec)
}
