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

var (
	teststrings = []string{
		"g",
		"e",
		"n",
	}
	emptystrings []string
	testapp      = &App{emptystrings, 0, 8, false}
)

func TestNewApp(t *testing.T) {
	type args struct {
		t   int
		val int
		s   bool
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
		{"newapp_test", args{0, 8, false}, testapp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApp(tt.args.t, tt.args.val, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_Generate(t *testing.T) {
	type fields struct {
		Generated []string
		AppType   APP_T
		IntVal    int
		Strong    bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"appgen_test", fields{teststrings, 0, 8, false}},
		{"appgen_test", fields{teststrings, 0, 3, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Generated: tt.fields.Generated,
				AppType:   tt.fields.AppType,
				IntVal:    tt.fields.IntVal,
				Strong:    tt.fields.Strong,
			}
			a.Generate()
		})
	}
}

func TestApp_String(t *testing.T) {
	type fields struct {
		Generated []string
		AppType   APP_T
		IntVal    int
		Strong    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"appstring_test", fields{teststrings, 0, 3, false}, "g" + "\n" + "e" + "\n" + "n" + "\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Generated: tt.fields.Generated,
				AppType:   tt.fields.AppType,
				IntVal:    tt.fields.IntVal,
				Strong:    tt.fields.Strong,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("App.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
