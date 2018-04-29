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
	"reflect"
	"testing"
)

const MAX_TESTS = 1000000

var (
	Toks = []*Token{
		{2, VOWEL},
		{0, CONSONANT},
		{0, VOWEL},
	}
	Newtoks_ubba = []*Token{
		{2, VOWEL},
		{0, CONSONANT},
		{0, CONSONANT},
		{0, VOWEL},
	}
	Newtoks_ubab = []*Token{
		{2, VOWEL},
		{0, CONSONANT},
		{0, VOWEL},
		{0, CONSONANT},
	}
	parsed_tokens = []*Token{}
)
var Tok = &Token{0, CONSONANT}

func Test_leet(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"leet_test", args{"A"}, "A"},
		{"leet_test", args{"E"}, "3"},
		{"leet_test", args{"S"}, "$"},
		{"leet_test", args{"O"}, "0"},
		{"leet_test", args{"T"}, "T"},
		{"leet_test", args{"F"}, "F"},
		{"leet_test", args{"I"}, "1"},
		{"leet_test", args{"B"}, "b"},
		{"leet_test", args{"G"}, "9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leet(tt.args.pass); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS*(len(Leet[tt.args.pass])-1); i++ {
					got = leet(tt.args.pass)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("leet() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestNewToken(t *testing.T) {
	token := &Token{2, CONSONANT}
	tests := []struct {
		name string
		want *Token
	}{
		{"new_tok", token},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewToken(); !reflect.DeepEqual(got, tt.want) {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got = NewToken()
					if reflect.DeepEqual(got, tt.want) {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("NewToken() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestToken_String(t *testing.T) {
	type fields struct {
		Token     int
		TokenType TOKEN_T
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test_token_string", fields{0, VOWEL}, "A"},
		{"test_token_string", fields{1, VOWEL}, "O"},
		{"test_token_string", fields{2, VOWEL}, "U"},
		{"test_token_string", fields{3, VOWEL}, "I"},
		{"test_token_string", fields{4, VOWEL}, "E"},
		{"test_token_string", fields{0, CONSONANT}, "B"},
		{"test_token_string", fields{1, CONSONANT}, "C"},
		{"test_token_string", fields{2, CONSONANT}, "D"},
		{"test_token_string", fields{3, CONSONANT}, "F"},
		{"test_token_string", fields{4, CONSONANT}, "G"},
		{"test_token_string", fields{5, CONSONANT}, "H"},
		{"test_token_string", fields{6, CONSONANT}, "J"},
		{"test_token_string", fields{7, CONSONANT}, "K"},
		{"test_token_string", fields{8, CONSONANT}, "L"},
		{"test_token_string", fields{9, CONSONANT}, "M"},
		{"test_token_string", fields{10, CONSONANT}, "N"},
		{"test_token_string", fields{11, CONSONANT}, "P"},
		{"test_token_string", fields{12, CONSONANT}, "Q"},
		{"test_token_string", fields{13, CONSONANT}, "R"},
		{"test_token_string", fields{14, CONSONANT}, "S"},
		{"test_token_string", fields{15, CONSONANT}, "T"},
		{"test_token_string", fields{16, CONSONANT}, "V"},
		{"test_token_string", fields{17, CONSONANT}, "W"},
		{"test_token_string", fields{18, CONSONANT}, "X"},
		{"test_token_string", fields{19, CONSONANT}, "Y"},
		{"test_token_string", fields{20, CONSONANT}, "Z"},
		{"test_token_string", fields{21, CONSONANT}, "*"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &Token{
				Token:     tt.fields.Token,
				TokenType: tt.fields.TokenType,
			}
			if got := tok.String(); got != tt.want {
				t.Errorf("Token.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSecret(t *testing.T) {
	secret := &Secret{}
	tests := []struct {
		name string
		want *Secret
	}{
		{"new_sec_test", secret},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSecret(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret_nextToken(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Token
	}{
		{"next_tok_test", fields{Toks, 3}, args{0}, &Token{0, CONSONANT}},
		{"next_tok_test", fields{Toks, 3}, args{1}, &Token{0, VOWEL}},
		{"next_tok_test", fields{Toks, 3}, args{2}, &Token{0, VOWEL}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.nextToken(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Secret.nextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret_insertToken(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	type args struct {
		token *Token
		num   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Token
	}{
		{"insert_test", fields{Toks, 3}, args{Tok, 2}, Newtoks_ubba},
		{"insert_test", fields{Toks, 3}, args{Tok, 3}, Newtoks_ubab},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.insertToken(tt.args.token, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Secret.insertToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret_generateTokens(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	type args struct {
		length int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"gentoken_test", fields{Toks, 3}, args{3}},
		{"gentoken_test", fields{Toks, 4}, args{6}},
		{"gentoken_test", fields{Toks, 7}, args{9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			s.generateTokens(tt.args.length)
		})
	}
}

func TestSecret_typeToken(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	type args struct {
		ttype TOKEN_T
		token *Token
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"typetoken_test", fields{Toks, 3}, args{CONSONANT, Tok}, true},
		{"typetoken_test", fields{Toks, 3}, args{DIPTHONG, Tok}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.typeToken(tt.args.ttype, tt.args.token); got != tt.want {
				t.Errorf("Secret.typeToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret_tune(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	type args struct {
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"tune_test", fields{Toks, 3}, args{"ABBA"}, "abb"},
		{"tune_test", fields{Toks, 3}, args{"ABBA"}, "Abb"},
		{"tune_test", fields{Toks, 3}, args{"ABBA"}, "aBb"},
		{"tune_test", fields{Toks, 3}, args{"ABBA"}, "abB"},
		{"tune_test", fields{Toks, 3}, args{"ABBA"}, "aBB"},
		{"tune_test", fields{Toks, 3}, args{"ABBA"}, "a66"},
		{"tune_test", fields{Toks, 3}, args{"ABBA"}, "*B6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.tune(tt.args.val); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got := s.tune(tt.args.val)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("Secret.tune() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSecret_parse(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.parse(); got != tt.want {
				t.Errorf("Secret.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret_PassWord(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	type args struct {
		length int
		strong bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.PassWord(tt.args.length, tt.args.strong); got != tt.want {
				check := false
				for i := 0; i < 10000; i++ {
					got := s.PassWord(tt.args.length, tt.args.strong)
					if got == tt.want {
						check = true
						break
					}
				}

				if check == true {
					t.Errorf("Secret.PassWord() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSecret_PassPhrase(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	type args struct {
		words  int
		strong bool
		random bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.PassPhrase(tt.args.words, tt.args.strong, tt.args.random); got != tt.want {
				t.Errorf("Secret.PassPhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret_String(t *testing.T) {
	type fields struct {
		Tokens []*Token
		Length int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"secstring_test", fields{Toks, 3}, "UBA"},
		{"secstring_test", fields{Toks, 3}, "uba"},
		{"secstring_test", fields{Toks, 3}, "U6A"},
		{"secstring_test", fields{Toks, 3}, "Uba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Secret{
				Tokens: tt.fields.Tokens,
				Length: tt.fields.Length,
			}
			if got := s.String(); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got := s.String()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("Secret.String() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
