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
	//	"errors"

	docopt "github.com/docopt/docopt-go"
	pw "github.com/hIMEI29A/pwgator/pwgator"
)

var version = "v0.1.3"

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

var (
	Configurator = NewConfig()
	PwArgs       = os.Args[1:]
)

type Config struct {
	Length  int
	Phrases int
	Strong  bool
	Random  bool
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) setArgs(args []string) error {
	var newerr error
	opts, err := docopt.ParseArgs(usage, args, version)

	if err != nil {
		newerr = err
	}

	switch {
	case opts["--phrase"]:
		if l := opts["LENGTH"]; l == nil {
			config.Length = 2
		} else {
			config.Length = pw.AToi(l.(string))
		}

		config.Phrases = 1
		config.Strong = opts["--strong"].(bool)
		config.Random = opts["-r"].(bool)

	default:
		if l := opts["LENGTH"]; l == nil {
			config.Length = 8
		} else {
			config.Length = pw.AToi(l.(string))
		}

		config.Phrases = 0
		config.Strong = opts["--strong"].(bool)
	}

	return newerr
}

func main() {
	if err := Configurator.setArgs(PwArgs); err != nil {
		pw.ErrFatal(err)
	}

	app := pw.NewApp(Configurator.Phrases, Configurator.Length, Configurator.Strong, Configurator.Random)
	app.Generate()

	fmt.Println(app.String())
}
