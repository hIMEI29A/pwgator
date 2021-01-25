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

var version = "v0.1.3"

var usage = `pwgator - humanized passphrases generator.

Usage:
  pwgator [-h | --help] | [--version]
  pwgator [LENGTH] [-s | --strong] [-p | --phrase [-r]]

Arguments:
  LENGTH               Secret's length (OPTIONAL). Default is 8 for password, 2 for phrase.

Options:
  -h --help            Show this screen.
  --version            Show version.
  -s --strong          Non-humanized generation.
  -p --phrase          Generate passphrases (with random template).
  -r                   Generate random passphrases.

`

var (
	configurator = NewConfig()
	pwArgs       = os.Args[1:]
)

type config struct {
	length  int
	phrases int
	strong  bool
	random  bool
}

func NewConfig() *config {
	return &config{}
}

func (c *config) setArgs(args []string) error {
	var newerr error
	opts, err := docopt.ParseArgs(usage, args, version)

	if err != nil {
		newerr = err
	}

	switch {
	case opts["--phrase"]:
		if l := opts["LENGTH"]; l == nil {
			c.length = 2
		} else {
			c.length = pw.AToi(l.(string))
		}

		c.phrases = 1
		c.strong = opts["--strong"].(bool)
		c.random = opts["-r"].(bool)

	default:
		if l := opts["LENGTH"]; l == nil {
			c.length = 8
		} else {
			c.length = pw.AToi(l.(string))
		}

		c.phrases = 0
		c.strong = opts["--strong"].(bool)
	}

	return newerr
}

func main() {
	if err := configurator.setArgs(pwArgs); err != nil {
		pw.ErrFatal(err)
	}

	app := pw.NewApp(configurator.phrases, configurator.length, configurator.strong, configurator.random)
	app.Generate()

	fmt.Println(app.GetString())
}
