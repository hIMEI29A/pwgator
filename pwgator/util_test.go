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

import "testing"

func Test_ranint(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ranint_test", args{10}, 0},
		{"ranint_test", args{10}, 1},
		{"ranint_test", args{10}, 2},
		{"ranint_test", args{10}, 3},
		{"ranint_test", args{10}, 4},
		{"ranint_test", args{10}, 5},
		{"ranint_test", args{10}, 6},
		{"ranint_test", args{10}, 7},
		{"ranint_test", args{10}, 8},
		{"ranint_test", args{10}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ranint(tt.args.num); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got := ranint(tt.args.num)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("ranint() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_ran(t *testing.T) {
	type args struct {
		num int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"ranint_test", args{int32(10)}, int32(0)},
		{"ranint_test", args{int32(10)}, int32(1)},
		{"ranint_test", args{int32(10)}, int32(2)},
		{"ranint_test", args{int32(10)}, int32(3)},
		{"ranint_test", args{int32(10)}, int32(4)},
		{"ranint_test", args{int32(10)}, int32(5)},
		{"ranint_test", args{int32(10)}, int32(6)},
		{"ranint_test", args{int32(10)}, int32(7)},
		{"ranint_test", args{int32(10)}, int32(8)},
		{"ranint_test", args{int32(10)}, int32(9)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ran(tt.args.num); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got := ran(tt.args.num)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("ran() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_cran(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		{"cran_test", 33},
		{"cran_test", 34},
		{"cran_test", 35},
		{"cran_test", 36},
		{"cran_test", 37},
		{"cran_test", 38},
		{"cran_test", 39},
		{"cran_test", 40},
		{"cran_test", 42},
		{"cran_test", 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cran(); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got := cran()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("cran() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestErrFatal(t *testing.T) {
	var testerror error
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{"e", args{testerror}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrFatal(tt.args.err)
		})
	}
}

func TestAToi(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"atoi_test", args{"10"}, 10},
		{"atoi_test", args{"1000000"}, 1000000},
		{"atoi_test", args{"2308756"}, 2308756},
		{"atoi_test", args{"0"}, 0},
		{"atoi_test", args{"-10"}, -10},
		{"atoi_test", args{"-2308756"}, -2308756},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AToi(tt.args.a); got != tt.want {
				t.Errorf("AToi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genstrong(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"genstrong_test", args{1}, "@"},
		{"genstrong_test", args{1}, "a"},
		{"genstrong_test", args{1}, "b"},
		{"genstrong_test", args{1}, "c"},
		{"genstrong_test", args{1}, "C"},
		{"genstrong_test", args{1}, "B"},
		{"genstrong_test", args{1}, "1"},
		{"genstrong_test", args{1}, ","},
		{"genstrong_test", args{1}, "y"},
		{"genstrong_test", args{1}, "&"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genstrong(tt.args.length); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got = genstrong(tt.args.length)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("genstrong() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_coin(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"coin_test", true},
		{"coin_test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coin(); got != tt.want {
				check := false
				for i := 0; i < 10; i++ {
					got := coin()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("coin() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_cointh(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"coin_test", true},
		{"coin_test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cointh(); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS; i++ {
					got := coin()
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("cointh() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_down(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"down_test", args{"BAkAbaKA"}, "bakabaka"},
		{"down_test", args{"qwerty"}, "qwerty"},
		{"down_test", args{"BABAK"}, "babak"},
		{"down_test", args{"QWERTYPOASDFG"}, "qwertypoasdfg"},
		{"down_test", args{"A"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := down(tt.args.value); got != tt.want {
				t.Errorf("down() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downize(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"downize_test", args{"ABAK"}, "abak"},
		{"downize_test", args{"ABAK"}, "Abak"},
		{"downize_test", args{"ABAK"}, "aBak"},
		{"downize_test", args{"ABAK"}, "abAk"},
		{"downize_test", args{"ABAK"}, "abaK"},

		{"downize_test", args{"KABA"}, "kaba"},
		{"downize_test", args{"KABA"}, "Kaba"},
		{"downize_test", args{"KABA"}, "KAba"},
		{"downize_test", args{"KABA"}, "kaBa"},
		{"downize_test", args{"KABA"}, "kabA"},

		{"downize_test", args{"ABKA"}, "abka"},
		{"downize_test", args{"ABKA"}, "Abka"},
		{"downize_test", args{"ABKA"}, "aBka"},
		{"downize_test", args{"ABKA"}, "abKa"},
		{"downize_test", args{"ABKA"}, "abkA"},

		{"downize_test", args{"BAKA"}, "baka"},
		{"downize_test", args{"BAKA"}, "Baka"},
		{"downize_test", args{"BAKA"}, "bAka"},
		{"downize_test", args{"BAKA"}, "baKa"},
		{"downize_test", args{"BAKA"}, "bakA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := downize(tt.args.pass); got != tt.want {
				check := false
				for i := 0; i < MAX_TESTS*len(tt.args.pass); i++ {
					got = downize(tt.args.pass)
					if got == tt.want {
						check = true
						break
					}
				}

				if check != true {
					t.Errorf("downize() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
