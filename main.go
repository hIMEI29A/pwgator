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

	gen "github.com/hIMEI29A/pwghen/libpwghen"
	//"github.com/docopt/docopt-go"
)

func main() {
	secret := gen.NewSecret()
	fmt.Println(secret.PassWord(4))
	fmt.Println(secret.PassPhrase(3))
}

/*
var usage = `gopwgen - humanized passwords generator.

Usage:
  gopwgen [-h | --help] | [--version]
  gopwgen [-r] LENGTH

Options:
  -h --help            Show this screen.
  --version            Show version.
  --interactive        Launch interactive master
  -o --output          Save report to file
  -j                   JSONed output

`
*/
