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

/**
File passphrase.go contains variables and functions for passphrases
generating. Here is two modes for this: generating with template (default) and
random generation. In first the passhrase's length (in characters) is fixed.
In both modes generation is done with given words number.

Passphrase length by words number in template mode:

	2 words 12 characters
	3 words 16 characters
	4 words 19 characters

In random mode passhrase's length in chars is random.
*/

package pwgator

import (
	"reflect"
	"testing"
)

func TestGetPhraseTemplate(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPhraseTemplate(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPhraseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomTemplate(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomTemplate(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
