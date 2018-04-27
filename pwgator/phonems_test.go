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

func TestTOKEN_T_String(t *testing.T) {
	tests := []struct {
		name string
		t    TOKEN_T
		want string
	}{
		{"tok_t", VOWEL, "VOWEL"},
		{"tok_t", DIPTHONG, "DIPTHONG"},
		{"tok_t", CONSONANT, "CONSONANT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("TOKEN_T.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTOKEN_T_TokenNames(t *testing.T) {
	tests := []struct {
		name string
		t    TOKEN_T
		want []string
	}{
		{"t_names", VOWEL, Vowels},
		{"t_names", DIPTHONG, Diphtongs},
		{"t_names", CONSONANT, Consonants},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.TokenNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TOKEN_T.TokenNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDIPTHONG_T_String(t *testing.T) {
	tests := []struct {
		name string
		d    DIPTHONG_T
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
				t.Errorf("DIPTHONG_T.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diftong_t(t *testing.T) {
	tests := []struct {
		name string
		want DIPTHONG_T
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
			if got := diftong_t(); got != tt.want {
				check := false
				for i := 0; i < 10*(len(Diphtongs)-1); i++ {
					got = diftong_t()

					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("diftong_t() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestVOWEL_T_String(t *testing.T) {
	tests := []struct {
		name string
		v    VOWEL_T
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
				t.Errorf("VOWEL_T.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVowel_t(t *testing.T) {
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
			if got := Vowel_t(tt.args.char); got != tt.want {
				t.Errorf("Vowel_t() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vowel_t(t *testing.T) {
	tests := []struct {
		name string
		want VOWEL_T
	}{
		{"vow_rand", A},
		{"vow_rand", O},
		{"vow_rand", U},
		{"vow_rand", I},
		{"vow_rand", E},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowel_t(); got != tt.want {
				check := false
				for i := 0; i < 10*(len(Vowels)-1); i++ {
					got = vowel_t()

					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("vowel_t() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCONSONANT_T_String(t *testing.T) {
	tests := []struct {
		name string
		c    CONSONANT_T
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
				t.Errorf("CONSONANT_T.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsonant_t(t *testing.T) {
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
			if got := Consonant_t(tt.args.char); got != tt.want {
				t.Errorf("Consonant_t() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consonant_t(t *testing.T) {
	tests := []struct {
		name string
		want CONSONANT_T
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
			if got := consonant_t(); got != tt.want {
				check := false
				for i := 0; i < 10*(len(Consonants)-1); i++ {
					got = consonant_t()

					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("consonant_t() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
