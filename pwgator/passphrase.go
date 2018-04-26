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

const SPACE = " "

var twoWord = [][]int{
	{6, 5},
	{6, 6},
	{5, 6},
}

var threeWords = [][]int{
	{3, 5, 5},
	{6, 3, 5},
	{5, 3, 6},
	{6, 3, 6},
	{6, 6, 3},
	{3, 3, 6},
}

var fourWords = [][]int{
	{3, 2, 5, 5},
	{5, 3, 2, 5},
}

var phrases = map[int][][]int{
	2: twoWord,
	3: threeWords,
	4: fourWords,
}

func GetPhraseTemplate(num int) []int {
	var template []int

	if templates, ok := phrases[num]; ok {
		r := ranint(len(templates))
		template = templates[r]
	}

	return template
}
