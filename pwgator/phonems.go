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
	"fmt"
)

// TOKEN_T

type TOKEN_T int

const (
	VOWEL TOKEN_T = iota
	DIPTHONG
	CONSONANT
)

var Tokens = [...]string{
	"VOWEL",
	"DIPTHONG",
	"CONSONANT",
}

func (t TOKEN_T) String() string {
	return fmt.Sprintf("%s", Tokens[t])
}

func (t TOKEN_T) TokenNames() []string {
	var names []string

	switch t {
	case CONSONANT:
		names = Consonants

	case DIPTHONG:
		names = Diphtongs

	case VOWEL:
		names = Vowels
	}

	return names
}

// DIPHTONG_T

type DIPTHONG_T int

const (
	AE DIPTHONG_T = iota
	AH
	CH
	EE
	EI
	GH
	IE
	NG
	KA
	AO
	DZ
	MB
	MH
	OI
	OA
	OH
	OO
	PH
	QU
	SH
	TH
)

var Diphtongs = []string{
	"AE",
	"AH",
	"CH",
	"EE",
	"EI",
	"GH",
	"IE",
	"NG",
	"KA",
	"AO",
	"DZ",
	"MB",
	"MH",
	"OI",
	"OA",
	"OH",
	"OO",
	"PH",
	"QU",
	"SH",
	"TH",
}

func (d DIPTHONG_T) String() string {
	return fmt.Sprintf("%s", Diphtongs[d])
}

func diftong_t() DIPTHONG_T {
	return DIPTHONG_T(ranint(len(Diphtongs)))
}

// VOWEL_T

type VOWEL_T int

const (
	A VOWEL_T = iota
	O
	U
	I
	E
)

var Vowels = []string{
	"A",
	"O",
	"U",
	"I",
	"E",
}

func (v VOWEL_T) String() string {
	return fmt.Sprintf("%s", Vowels[v])
}

func Vowel_t(char string) bool {
	check := false

	for i := range Vowels {
		if char == Vowels[i] {
			check = true
			break
		}
	}

	return check
}

func vowel_t() VOWEL_T {
	return VOWEL_T(ranint(len(Vowels)))
}

// CONSONANT_T

type CONSONANT_T int

const (
	B CONSONANT_T = iota
	C
	D
	F
	G
	H
	J
	K
	L
	M
	N
	P
	Q
	R
	S
	T
	V
	W
	X
	Y
	Z
	ZZ // special for asterisk "*" as random delim
)

var Consonants = []string{
	"B",
	"C",
	"D",
	"F",
	"G",
	"H",
	"J",
	"K",
	"L",
	"M",
	"N",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"*",
}

func (c CONSONANT_T) String() string {
	return fmt.Sprintf("%s", Consonants[c])
}

func Consonant_t(char string) bool {
	check := false

	for i := range Consonants {
		if char == Consonants[i] {
			check = true
			break
		}
	}

	return check
}

func consonant_t() CONSONANT_T {
	return CONSONANT_T(ranint(len(Consonants)))
}
