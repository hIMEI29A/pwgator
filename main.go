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

	docopt "github.com/docopt/docopt-go"
	pw "github.com/hIMEI29A/pwgator/pwgator"
)

var version = "v0.1.2"

var usage = `pwgator - humanized passphrases generator.

Usage:
  pwgator [-h | --help] | [--version]
  pwgator [LENGTH] [-s | --strong] [-p | --phrase [-r]]

Arguments:
  LENGTH               Secret's length (OPTIONAL). Default is 8 for word, 2 for phrase.

Options:
  -h --help            Show this screen.
  --version            Show version.
  -s --strong          Non-humanized generation.
  -p --phrase          Generate passphrases (with random template).
  -r                   Generate random passphrases.

`

func main() {
	opts, _ := docopt.ParseArgs(usage, os.Args[1:], version)

	var (
		length  int
		phrases int
		strong  bool
		random  bool
	)

	switch {
	case opts["--phrase"]:
		if l := opts["LENGTH"]; l == nil {
			length = 2
		} else {
			length = pw.AToi(l.(string))
		}

		phrases = 1
		strong = opts["--strong"].(bool)
		random = opts["-r"].(bool)

	default:
		if l := opts["LENGTH"]; l == nil {
			length = 8
		} else {
			length = pw.AToi(l.(string))
		}

		phrases = 0
		strong = opts["--strong"].(bool)

	}
	app := pw.NewApp(phrases, length, strong, random)
	app.Generate()

	fmt.Println(app.String())
}
