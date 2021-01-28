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

const maxTests = 1000000

var (
	Toks = []*token{
		{2, vowel},
		{0, consonant},
		{0, vowel},
	}
	Newtoks_ubba = []*token{
		{2, vowel},
		{0, consonant},
		{0, consonant},
		{0, vowel},
	}
	Newtoks_ubab = []*token{
		{2, vowel},
		{0, consonant},
		{0, vowel},
		{0, consonant},
	}
	parsed_tokens = []*token{}
)
var Tok = &token{0, consonant}

func Test_makeLeet(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"makeLeet_test", args{"A"}, "A"},
		{"makeLeet_test", args{"E"}, "3"},
		{"makeLeet_test", args{"S"}, "$"},
		{"makeLeet_test", args{"O"}, "0"},
		{"makeLeet_test", args{"T"}, "T"},
		{"makeLeet_test", args{"F"}, "F"},
		{"makeLeet_test", args{"I"}, "1"},
		{"makeLeet_test", args{"B"}, "b"},
		{"makeLeet_test", args{"G"}, "9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeLeet(tt.args.pass); got != tt.want {
				check := false
				for i := 0; i < maxTests*(len(leet[tt.args.pass])-1); i++ {
					got = makeLeet(tt.args.pass)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("makeLeet() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestNewToken(t *testing.T) {
	to := &token{2, consonant}
	tests := []struct {
		name string
		want *token
	}{
		{"new_tok", to},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newToken(); !reflect.DeepEqual(got, tt.want) {
				check := false
				for i := 0; i < maxTests; i++ {
					got = newToken()
					if reflect.DeepEqual(got, tt.want) {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("newToken() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestToken_String(t *testing.T) {
	type fields struct {
		to        int
		tokenType tokenType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test_token_string", fields{0, vowel}, "A"},
		{"test_token_string", fields{1, vowel}, "O"},
		{"test_token_string", fields{2, vowel}, "U"},
		{"test_token_string", fields{3, vowel}, "I"},
		{"test_token_string", fields{4, vowel}, "E"},
		{"test_token_string", fields{0, consonant}, "B"},
		{"test_token_string", fields{1, consonant}, "C"},
		{"test_token_string", fields{2, consonant}, "D"},
		{"test_token_string", fields{3, consonant}, "F"},
		{"test_token_string", fields{4, consonant}, "G"},
		{"test_token_string", fields{5, consonant}, "H"},
		{"test_token_string", fields{6, consonant}, "J"},
		{"test_token_string", fields{7, consonant}, "K"},
		{"test_token_string", fields{8, consonant}, "L"},
		{"test_token_string", fields{9, consonant}, "M"},
		{"test_token_string", fields{10, consonant}, "N"},
		{"test_token_string", fields{11, consonant}, "P"},
		{"test_token_string", fields{12, consonant}, "Q"},
		{"test_token_string", fields{13, consonant}, "R"},
		{"test_token_string", fields{14, consonant}, "S"},
		{"test_token_string", fields{15, consonant}, "T"},
		{"test_token_string", fields{16, consonant}, "V"},
		{"test_token_string", fields{17, consonant}, "W"},
		{"test_token_string", fields{18, consonant}, "X"},
		{"test_token_string", fields{19, consonant}, "Y"},
		{"test_token_string", fields{20, consonant}, "Z"},
		{"test_token_string", fields{21, consonant}, "*"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &token{
				token:     tt.fields.to,
				tokenType: tt.fields.tokenType,
			}
			if got := tok.String(); got != tt.want {
				t.Errorf("token.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSecret(t *testing.T) {
	s := &secret{}
	tests := []struct {
		name string
		want *secret
	}{
		{"new_sec_test", s},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSecret(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret_nextToken(t *testing.T) {
	type fields struct {
		tokens []*token
		length int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *token
	}{
		{"next_tok_test", fields{Toks, 3}, args{0}, &token{0, consonant}},
		{"next_tok_test", fields{Toks, 3}, args{1}, &token{0, vowel}},
		{"next_tok_test", fields{Toks, 3}, args{2}, &token{0, vowel}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.length,
			}
			if got := s.nextToken(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("secret.nextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret_insertToken(t *testing.T) {
	type fields struct {
		tokens []*token
		length int
	}
	type args struct {
		to  *token
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*token
	}{
		{"insert_test", fields{Toks, 3}, args{Tok, 2}, Newtoks_ubba},
		{"insert_test", fields{Toks, 3}, args{Tok, 3}, Newtoks_ubab},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.length,
			}
			if got := s.insertToken(tt.args.to, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Secret.insertToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret_generateTokens(t *testing.T) {
	type fields struct {
		tokens []*token
		l      int
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
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.l,
			}
			s.generateTokens(tt.args.length)
		})
	}
}

func Test_secret_typeToken(t *testing.T) {
	type fields struct {
		to     []*token
		length int
	}
	type args struct {
		ttype tokenType
		token *token
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"typetoken_test", fields{Toks, 3}, args{consonant, Tok}, true},
		{"typetoken_test", fields{Toks, 3}, args{diphthong, Tok}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &secret{
				tokens: tt.fields.to,
				length: tt.fields.length,
			}
			if got := s.typeToken(tt.args.ttype, tt.args.token); got != tt.want {
				t.Errorf("secret.typeToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret_tune(t *testing.T) {
	type fields struct {
		tokens []*token
		length int
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
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.length,
			}
			if got := s.tune(tt.args.val); got != tt.want {
				check := false
				for i := 0; i < maxTests; i++ {
					got := s.tune(tt.args.val)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("secret.tune() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_secret_parse(t *testing.T) {
	type fields struct {
		tokens []*token
		length int
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
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.length,
			}
			if got := s.parse(); got != tt.want {
				t.Errorf("secret.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret_passWord(t *testing.T) {
	type fields struct {
		tokens []*token
		length int
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
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.length,
			}
			if got := s.passWord(tt.args.length, tt.args.strong); got != tt.want {
				check := false
				for i := 0; i < 10000; i++ {
					got := s.passWord(tt.args.length, tt.args.strong)
					if got == tt.want {
						check = true
						break
					}
				}

				if check == true {
					t.Errorf("secret.passWord() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_secret_PassPhrase(t *testing.T) {
	type fields struct {
		tokens []*token
		length int
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
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.length,
			}
			if got := s.passPhrase(tt.args.words, tt.args.strong, tt.args.random); got != tt.want {
				t.Errorf("secret.passPhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret_String(t *testing.T) {
	type fields struct {
		tokens []*token
		length int
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
			s := &secret{
				tokens: tt.fields.tokens,
				length: tt.fields.length,
			}
			if got := s.String(); got != tt.want {
				check := false
				for i := 0; i < maxTests; i++ {
					got := s.String()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("secret.String() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
