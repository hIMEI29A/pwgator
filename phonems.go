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

type tokenType int

const (
	vowel tokenType = iota
	diphthong
	consonant
)

var tokens = [...]string{
	"vowel",
	"diphthong",
	"consonant",
}

// String stringizes tokenType's
func (t tokenType) String() string {
	return fmt.Sprintf("%s", tokens[t])
}

func (t tokenType) tokenNames() []string {
	var names []string

	switch t {
	case consonant:
		names = consonants

	case diphthong:
		names = diphthongs

	case vowel:
		names = vowels
	}

	return names
}

type diphthongType int

const (
	AE diphthongType = iota
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

var diphthongs = []string{
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

// String stringizes diphthongType's
func (d diphthongType) String() string {
	return fmt.Sprintf("%s", diphthongs[d])
}

func getDiphthong() diphthongType {
	return diphthongType(ranint(len(diphthongs)))
}

type vowelType int

const (
	A vowelType = iota
	O
	U
	I
	E
)

var vowels = []string{
	"A",
	"O",
	"U",
	"I",
	"E",
}

// String stringizes vowelType's
func (v vowelType) String() string {
	return fmt.Sprintf("%s", vowels[v])
}

func isVowel(char string) bool {
	check := false

	for i := range vowels {
		if char == vowels[i] {
			check = true
			break
		}
	}

	return check
}

func getVowel() vowelType {
	return vowelType(ranint(len(vowels)))
}

type consonantType int

const (
	B consonantType = iota
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

var consonants = []string{
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

// String stringizes consonantType's
func (c consonantType) String() string {
	return fmt.Sprintf("%s", consonants[c])
}

func isConsonant(char string) bool {
	check := false

	for i := range consonants {
		if char == consonants[i] {
			check = true
			break
		}
	}

	return check
}

func getConsonant() consonantType {
	return consonantType(ranint(len(consonants)))
}
