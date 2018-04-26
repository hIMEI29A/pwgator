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
	gator "github.com/hIMEI29A/pwgator/pwgator"
)

var version = "v0.0.1-alpha"

var usage = `pwgator - humanized passphrases generator.

Usage:
  pwgator [-h | --help] | [--version]
  pwgator [LENGTH] [-s | --strong] [-p | --phrase]

Arguments:
  LENGTH               Secret length (OPTIONAL). Default is 8 for word, 2 for phrase.

Options:
  -h --help            Show this screen.
  --version            Show version.
  -s --strong          Non-humanized genaration.
  -p --phrases         Generate passphrases.

`

func main() {

	fmt.Println(os.Args[1:])
	opts, _ := docopt.ParseArgs(usage, os.Args[1:], version)

	var app *gator.App

	var conf struct {
		Length  int
		Phrases bool
		Strong  bool
	}

	opts.Bind(&conf)

	fmt.Println(conf.Length)

	if conf.Phrases && conf.Length == 0 {
		app = gator.NewApp(1, 2, conf.Strong)
	}

	if conf.Phrases && conf.Length != 0 {
		app = gator.NewApp(1, conf.Length, conf.Strong)
	}

	if !conf.Phrases && conf.Length == 0 {
		app = gator.NewApp(0, 8, conf.Strong)
	}

	if !conf.Phrases && conf.Length != 0 {
		app = gator.NewApp(0, conf.Length, conf.Strong)
	}

	fmt.Println(conf.Length)

	app.Generate()
	fmt.Println(app.String())

}

/*
	app := gator.NewApp(0, 8, false)

	app.Generate()

	fmt.Println(app.String())
}*/
