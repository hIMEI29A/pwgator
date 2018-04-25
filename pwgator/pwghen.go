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
	//"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const COIN_MAX = 100000

const (
	BASE_S int32 = 32
	BASE_E       = 128
)

var Leet = map[string][]string{
	"A":  {"a", "@", "*", "A"},
	"E":  {"e", "3", "E"},
	"S":  {"s", "$", "5", "S"},
	"O":  {"o", "0", "O"},
	"EE": {"ee", "33"},
	"T":  {"7", "t", "T"},
	"I":  {"i", "!", "1", "I"},
	"B":  {"b", "8", "6", "B"},
	"G":  {"g", "9", "G"},
}

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

func ranint(num int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	random := r.Intn(num)

	return random
}

func ran(num int32) int32 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	random := r.Int31n(num)

	return random
}

func cran() int32 {
	cRan := ran(BASE_E)

	if cRan > BASE_S {
		return cRan
	} else {
		cRan = cran()
	}
	return cRan
}

func ErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func genstrong(length int) string {
	var chars string

	for j := 0; j < length; j++ {
		char := strconv.QuoteRuneToASCII(cran())
		uchar, err := strconv.Unquote(char)
		ErrFatal(err)
		chars = chars + uchar
	}

	return chars
}

func coin() bool {
	value := false

	r := ranint(COIN_MAX)

	if r >= COIN_MAX/2 {
		value = true
	}

	return value
}

func cointh() bool {
	value := false

	r := ranint(COIN_MAX)

	if r%3 == 0 {
		value = true
	}

	return value
}

func down(value string) string {
	var s string
	for i := range value {
		s = s + string(unicode.ToLower(rune(value[i])))
	}

	return s
}

func downize(pass string) string {
	var newpass string

	for i := range pass {
		if c := coin(); c {
			newpass = newpass + down(string(pass[i]))
		} else {
			newpass = newpass + string(pass[i])
		}
	}

	return newpass
}

type Token struct {
	Token     int
	TokenType TOKEN_T
}

func NewToken() *Token {
	token := &Token{}

	token.TokenType = TOKEN_T(ranint(len(Tokens)))
	token.Token = ranint(len(token.TokenType.TokenNames()))

	return token
}

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

type Secret struct {
	Tokens []*Token
	Length int
}

func NewSecret() *Secret {
	secret := &Secret{}
	//secret.Length = length

	return secret
}

func (s *Secret) TypeToken(ttype TOKEN_T, token *Token) bool {
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

func (s *Secret) String() string {
	var sec string

	for i := range s.Tokens {
		sec = sec + s.Tokens[i].String()
	}

	return s.Tune(sec)
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

func (s *Secret) Tune(val string) string {
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

func (s *Secret) Parse() string {
	//var word string
	newtokens := s.Tokens

	for i := range newtokens {

		if s.TypeToken(VOWEL, newtokens[i]) == true && s.TypeToken(VOWEL, s.nextToken(i)) == true {
			newtoken := NewToken()
			newtoken.Token = int(consonant_t())
			newtoken.TokenType = CONSONANT

			newtokens = s.insertToken(newtoken, i+1)
		}

		if s.TypeToken(CONSONANT, newtokens[i]) == true && s.TypeToken(CONSONANT, s.nextToken(i)) == true {
			newtoken := NewToken()
			newtoken.Token = int(vowel_t())
			newtoken.TokenType = VOWEL

			newtokens = s.insertToken(newtoken, i+1)
		}
	}

	s.Tokens = newtokens

	return s.String()
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

func (s *Secret) PassWord(length int) string {
	s.generateTokens(length)
	return s.Parse()
}

func (s *Secret) PassPhrase(words int) string {
	template := GetPhraseTemplate(words)
	var phrase string

	for i := range template {
		phrase = phrase + SPACE + s.PassWord(template[i])
	}

	return strings.TrimPrefix(phrase, SPACE)
}
