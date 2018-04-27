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
	"log"
	"math/rand"
	"strconv"
	"time"
	"unicode"
)

const COIN_MAX = 10000

const (
	BASE_S int32 = 32
	BASE_E       = 128
)

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

func AToi(a string) int {
	r, err := strconv.Atoi(a)
	ErrFatal(err)

	return r
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
