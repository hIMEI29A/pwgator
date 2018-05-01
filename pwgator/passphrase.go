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

//File passphrase.go contains variables and functions for passphrases
//generating. Here is two modes for this: generating with template (default) and
//random generation. In first the passhrase's length (in characters) is fixed.
//In both modes generation is done with given words number.
//
//Passphrase length by words number in template mode:
//
//	2 words 12 characters
//	3 words 16 characters
//	4 words 19 characters
//
//In random mode passhrase's length in chars is random.
package pwgator

const SPACE = " "

const (
	LETTERS_MIN = 2
	LETTERS_MAX = 8
)

var twoWord = [][]int{
	{6, 5},
	{5, 6},
}

var threeWords = [][]int{
	{3, 5, 6},
	{6, 3, 5},
	{5, 3, 6},
	{6, 5, 3},
	{5, 6, 3},
}

var fourWords = [][]int{
	{3, 5, 5, 3},
	{5, 5, 3, 3},
	{3, 3, 5, 5},
	{5, 3, 3, 5},
}

var phrases = map[int][][]int{
	2: twoWord,
	3: threeWords,
	4: fourWords,
}

// GetPhraseTemplate returns one of passphrase template's with given length.
func GetPhraseTemplate(num int) []int {
	var template []int

	if templates, ok := phrases[num]; ok {
		r := ranint(len(templates))
		template = templates[r]
	}

	return template
}

// GetRandomTemplate generates random passphrase's template
func GetRandomTemplate(num int) []int {
	var (
		template []int
		temp     int
	)

	for i := 0; i < num; i++ {
		temp = ranint(LETTERS_MAX)

		if temp < LETTERS_MIN {
			temp = ranint(LETTERS_MAX)
		}

		template = append(template, temp)
	}

	return template
}
