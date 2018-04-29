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

/*
Description

Pwgator - Golang package and application for humanized passwords and passphrases
generation. Inspired by  [pwgen](https://github.com/tytso/pwgen).

Usage

Call **pwgator** with `-h` or `--help` option to get full usage instruction.

	$ pwgator -h

	pwgator - humanized passphrases generator.

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


To generate pass words use

	pwgator

With given length:

	pwgator 666

Passphrases:

	pwgator -p

Random passphrases with 6 words length:

	pwgator -p -r 6

In strong mode

	pwgator -s


*/
package main
