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

// TOKEN_T is an enum for token types
type TOKEN_T int

// Enumerated token types
const (
	VOWEL TOKEN_T = iota
	DIPTHONG
	CONSONANT
)

// Tokens holds string representations of TOKEN_T
var Tokens = [...]string{
	"VOWEL",
	"DIPTHONG",
	"CONSONANT",
}

// String stringizes TOKEN_T's
func (t TOKEN_T) String() string {
	return fmt.Sprintf("%s", Tokens[t])
}

// TokenNames returns string values containing variable for each token type:
// Consonants for CONSONANT, Diphtongs for DIPTHONG, Vowels for VOWEL
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

// DIPHTHONG_T is an enum for two letter string literals
type DIPTHONG_T int

// Enumerated two letter string literals
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

// Diphtongs holds string representations of DIPTHONG_T
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

// String stringizes DIPTHONG_T's
func (d DIPTHONG_T) String() string {
	return fmt.Sprintf("%s", Diphtongs[d])
}

// diftong_t returns random DIPTHONG_T
func diftong_t() DIPTHONG_T {
	return DIPTHONG_T(ranint(len(Diphtongs)))
}

// VOWEL_T is an enum for vowels representing string literals
type VOWEL_T int

// Enumerated vowels representing string literals
const (
	A VOWEL_T = iota
	O
	U
	I
	E
)

// Vowels holds string representations of VOWEL_T
var Vowels = []string{
	"A",
	"O",
	"U",
	"I",
	"E",
}

// String stringizes VOWEL_T's
func (v VOWEL_T) String() string {
	return fmt.Sprintf("%s", Vowels[v])
}

// Vowel_t checks if given string is a string representation of VOWEL_T
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

// vowel_t returns random VOWEL_T
func vowel_t() VOWEL_T {
	return VOWEL_T(ranint(len(Vowels)))
}

// CONSONANT_T is an enum for consonants representing string literals
type CONSONANT_T int

// Enumerated consonants representing string literals
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

// Consonants holds string representations of CONSONANT_T
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

// String stringizes CONSONANT_T's
func (c CONSONANT_T) String() string {
	return fmt.Sprintf("%s", Consonants[c])
}

// Consonant_t checks if given string is a string representation of CONSONANT_T
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

// consonant_t returns random CONSONANT_T
func consonant_t() CONSONANT_T {
	return CONSONANT_T(ranint(len(Consonants)))
}
