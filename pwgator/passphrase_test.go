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

func TestGetPhraseTemplate(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	//{"template_rand", args{2}},
	//{"template_rand", args{3}},
	//{"template_rand", args{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPhraseTemplate(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPhraseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
