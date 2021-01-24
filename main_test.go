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

package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	pw "github.com/hIMEI29A/pwgator/pwgator"
)

var TestArgs = []string{
	"4",
	"-s",
	"-p",
	"-r",
}

func TestNewConfig(t *testing.T) {
	c := &config{}

	tests := []struct {
		name string
		want *config
	}{
		{"newconfig_test", c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_setArgs(t *testing.T) {
	type fields struct {
		length  int
		phrases int
		strong  bool
		random  bool
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"setargs_test", fields{0, 0, false, false}, args{TestArgs}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				length:  tt.fields.Length,
				phrases: tt.fields.Phrases,
				strong:  tt.fields.Strong,
				random:  tt.fields.Random,
			}
			if err := config.setArgs(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Config.setArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	os.Args = TestArgs
	configurator = NewConfig()

	tests := []struct {
		name string
	}{
		{"main_t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := configurator.setArgs(pwArgs); err != nil {
				t.Errorf("Arguments parsing error: %v", err)
			}

			app := pw.NewApp(
				configurator.phrases,
				configurator.length,
				configurator.strong,
				configurator.random,
			)

			if app == nil {
				t.Errorf("Config error")
			}

			app.Generate()

			if len(app.Generated) != 100 {
				t.Errorf("String generation's error")
			}

			if _, err := fmt.Println(app.String()); err != nil {
				t.Errorf("Printing error: %v", err)
			}
		})
	}
}
