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

func Test_tokenType_String(t *testing.T) {
	tests := []struct {
		name string
		t    tokenType
		want string
	}{
		{"tok_t", vowel, "vowel"},
		{"tok_t", diphthong, "diphthong"},
		{"tok_t", consonant, "consonant"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("tokenType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenType_tokenNames(t *testing.T) {
	tests := []struct {
		name string
		t    tokenType
		want []string
	}{
		{"t_names", vowel, vowels},
		{"t_names", diphthong, diphthongs},
		{"t_names", consonant, consonants},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.tokenNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenType.tokenNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diphthongType_String(t *testing.T) {
	tests := []struct {
		name string
		d    diphthongType
		want string
	}{
		{"dif_t", AE, "AE"},
		{"dif_t", AH, "AH"},
		{"dif_t", CH, "CH"},
		{"dif_t", EE, "EE"},
		{"dif_t", EI, "EI"},
		{"dif_t", GH, "GH"},
		{"dif_t", IE, "IE"},
		{"dif_t", NG, "NG"},
		{"dif_t", KA, "KA"},
		{"dif_t", AO, "AO"},
		{"dif_t", DZ, "DZ"},
		{"dif_t", MB, "MB"},
		{"dif_t", MH, "MH"},
		{"dif_t", OI, "OI"},
		{"dif_t", OA, "OA"},
		{"dif_t", OH, "OH"},
		{"dif_t", OO, "OO"},
		{"dif_t", PH, "PH"},
		{"dif_t", QU, "QU"},
		{"dif_t", SH, "SH"},
		{"dif_t", TH, "TH"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("diphthongType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDiphthong(t *testing.T) {
	tests := []struct {
		name string
		want diphthongType
	}{
		{"dif_rand", AE},
		{"dif_rand", AH},
		{"dif_rand", CH},
		{"dif_rand", EE},
		{"dif_rand", EI},
		{"dif_rand", GH},
		{"dif_rand", IE},
		{"dif_rand", NG},
		{"dif_rand", KA},
		{"dif_rand", AO},
		{"dif_rand", DZ},
		{"dif_rand", MB},
		{"dif_rand", MH},
		{"dif_rand", OI},
		{"dif_rand", OA},
		{"dif_rand", OH},
		{"dif_rand", OO},
		{"dif_rand", PH},
		{"dif_rand", QU},
		{"dif_rand", SH},
		{"dif_rand", TH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiphthong(); got != tt.want {
				check := false
				for i := 0; i < 10*(len(diphthongs)-1); i++ {
					got = getDiphthong()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("getDiphthong() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_vowelType_String(t *testing.T) {
	tests := []struct {
		name string
		v    vowelType
		want string
	}{
		{"vow_t", A, "A"},
		{"vow_t", O, "O"},
		{"vow_t", U, "U"},
		{"vow_t", I, "I"},
		{"vow_t", E, "E"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("vowelType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVowel(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"v", args{"A"}, true},
		{"v", args{"O"}, true},
		{"v", args{"R"}, false},
		{"v", args{"WE"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVowel(tt.args.char); got != tt.want {
				t.Errorf("isVowel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVowel(t *testing.T) {
	tests := []struct {
		name string
		want vowelType
	}{
		{"vow_rand", A},
		{"vow_rand", O},
		{"vow_rand", U},
		{"vow_rand", I},
		{"vow_rand", E},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVowel(); got != tt.want {
				check := false
				for i := 0; i < 10*(len(vowels)-1); i++ {
					got = getVowel()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("getVowel() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_consonantType_String(t *testing.T) {
	tests := []struct {
		name string
		c    consonantType
		want string
	}{
		{"cons_t", B, "B"},
		{"cons_t", C, "C"},
		{"cons_t", D, "D"},
		{"cons_t", F, "F"},
		{"cons_t", G, "G"},
		{"cons_t", H, "H"},
		{"cons_t", J, "J"},
		{"cons_t", K, "K"},
		{"cons_t", L, "L"},
		{"cons_t", M, "M"},
		{"cons_t", N, "N"},
		{"cons_t", P, "P"},
		{"cons_t", Q, "Q"},
		{"cons_t", R, "R"},
		{"cons_t", S, "S"},
		{"cons_t", T, "T"},
		{"cons_t", V, "V"},
		{"cons_t", W, "W"},
		{"cons_t", X, "X"},
		{"cons_t", Y, "Y"},
		{"cons_t", ZZ, "*"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("consonantType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isConsonant(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"vv", args{"A"}, false},
		{"vv", args{"O"}, false},
		{"vv", args{"R"}, true},
		{"vv", args{"WE"}, false},
		{"vv", args{"W"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isConsonant(tt.args.char); got != tt.want {
				t.Errorf("isConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConsonant(t *testing.T) {
	tests := []struct {
		name string
		want consonantType
	}{
		{"cons_rand", B},
		{"cons_rand", C},
		{"cons_rand", D},
		{"cons_rand", F},
		{"cons_rand", G},
		{"cons_rand", H},
		{"cons_rand", J},
		{"cons_rand", K},
		{"cons_rand", L},
		{"cons_rand", M},
		{"cons_rand", N},
		{"cons_rand", P},
		{"cons_rand", Q},
		{"cons_rand", R},
		{"cons_rand", S},
		{"cons_rand", T},
		{"cons_rand", V},
		{"cons_rand", W},
		{"cons_rand", X},
		{"cons_rand", Y},
		{"cons_rand", ZZ},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConsonant(); got != tt.want {
				check := false
				for i := 0; i < 10*(len(consonants)-1); i++ {
					got = getConsonant()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("getConsonant() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
