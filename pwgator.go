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

var leet = map[string][]string{
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

func makeLeet(pass string) string {
	var newpass string

	for i := range pass {
		if _, ok := leet[string(pass[i])]; ok {
			values := leet[string(pass[i])]
			trand := ranint(len(values))
			newpass = newpass + values[trand]
		} else {
			newpass = newpass + string(pass[i])
		}
	}

	return newpass
}

type token struct {
	token     int
	tokenType tokenType
}

// NewToken instantiates token type
func newToken() *token {
	t := &token{}

	t.tokenType = tokenType(ranint(len(tokens)))
	t.token = ranint(len(t.tokenType.tokenNames()))

	return t
}

// String gives string representation of token
func (t *token) String() string {
	var tokenized string

	switch t.tokenType {
	case consonant:
		tokenized = consonantType(t.token).String()

	case vowel:
		tokenized = vowelType(t.token).String()

	case diphthong:
		tokenized = diphthongType(t.token).String()

	}

	return tokenized
}

type secret struct {
	tokens []*token
	length int
}

func newSecret() *secret {
	s := &secret{}

	return s
}

func (s *secret) nextToken(num int) *token {
	t := &token{}

	if num < len(s.tokens)-1 {
		t = s.tokens[num+1]
	} else {
		t = s.tokens[len(s.tokens)-1]
	}
	return t
}

func (s *secret) insertToken(t *token, num int) []*token {
	newtoken := &token{}
	newtokens := s.tokens

	newtokens = append(newtokens, newtoken)
	copy(newtokens[num+1:], newtokens[num:])
	newtokens[num] = t

	return newtokens
}

func (s *secret) generateTokens(length int) {
	var st []*token

	for i := 0; i < length; i++ {
		t := newToken()
		st = append(st, t)
	}
	s.length = len(st)

	s.tokens = st
}

func (s *secret) typeToken(ttype tokenType, t *token) bool {
	check := false

	if t.tokenType == ttype {
		check = true
	}

	if t.tokenType == diphthong {
		tstring := t.String()

		if ttype == consonant {
			if isConsonant(string(tstring[0])) == true && isConsonant(string(tstring[1])) == true {
				check = true
			}
		}

		if ttype == vowel {
			if isVowel(string(tstring[0])) == true && isVowel(string(tstring[1])) == true {
				check = true
			}
		}
	}

	return check
}

func (s *secret) tune(val string) string {
	word := downize(makeLeet(val))
	var newword string

	if len(word) > s.length {
		for i := 0; i < s.length; i++ {
			newword = newword + string(word[i])
		}
	} else {
		newword = word
	}

	return newword
}

// parse makes manipulation with Secret's []*token
func (s *secret) parse() string {
	newtokens := s.tokens

	for i := range newtokens {
		// if vowel or diphthong with two vowels inside is followed by vowel,
		// or any diphthong is followed by vowel, replace this last vowel with new token
		if (s.typeToken(vowel, newtokens[i]) == true && s.typeToken(vowel, s.nextToken(i)) == true) ||
			(s.typeToken(diphthong, newtokens[i]) == true && s.typeToken(vowel, s.nextToken(i)) == true) {
			newtoken := newToken()
			newtoken.token = int(getConsonant())
			newtoken.tokenType = consonant

			newtokens = s.insertToken(newtoken, i+1)
		}

		// if consonant or diphthong with two consonant inside is followed by consonant,
		// replace this last consonant with new token
		if s.typeToken(consonant, newtokens[i]) == true && s.typeToken(consonant, s.nextToken(i)) == true {
			newtoken := newToken()
			newtoken.token = int(getVowel())
			newtoken.tokenType = vowel

			newtokens = s.insertToken(newtoken, i+1)
		}
	}

	s.tokens = newtokens

	return s.String()
}

func (s *secret) passWord(length int, strong bool) string {
	var p string

	if !strong {
		s.generateTokens(length)
		p = s.parse()
	} else {
		p = genstrong(length)
	}

	return p
}

func (s *secret) passPhrase(words int, strong, random bool) string {
	var (
		phrase   string
		template []int
	)

	template = getPhraseTemplate(words)

	if random == true {
		template = getRandomTemplate(words)
	}

	if !strong {
		for i := range template {
			phrase = phrase + " " + s.passWord(template[i], strong)
		}
	} else {
		for i := range template {
			phrase = phrase + " " + genstrong(template[i])
		}
	}

	return strings.TrimPrefix(phrase, " ")
}

// String stringizes Secret's []*token
func (s *secret) String() string {
	var sec string

	for i := range s.tokens {
		sec = sec + s.tokens[i].String()
	}

	return s.tune(sec)
}
