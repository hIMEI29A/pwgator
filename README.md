# pwgator

**pwgator** - Golang package and application for humanized passwords and passphrases 
generation. Inspired by great [pwgen](https://github.com/tytso/pwgen).

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/pwgator)](https://goreportcard.com/report/github.com/hIMEI29A/pwgator) [![GoDoc](https://godoc.org/github.com/hIMEI29A/pwgator?status.svg)](http://godoc.org/github.com/hIMEI29A/pwgator) [![Apache-2.0 License](https://img.shields.io/badge/license-Apache--2.0-red.svg)](LICENSE)

## TOC
- [Features](#features)
- [Version](#version)
- [Install](#install)
- [Usage](#usage)

## Features

- Random passwords and passphrases generation. It can be done in **humanized** and **strong** (fully random and with all special symbols) mode.

- Default pasword length is **8 symbols** (without **LENGTH** argument). Default passphrase's template is **2 words**.

- Default generation mode is **humanized**.

- Template driven passphrase generation. Templates available: 2 words, 3 words, 4 words. One template generated passphrases have different length.

- 100 concurrently generated passwords or passphrases output by one app call.

- Some leet-codes is used for humanisation random strings.

- Passwords:

	1. humanized
	
		Fo5h!po
		mH*nGon

		!*kE!YVZ
		o1tZ@*Po
		rUieIo@o
		
		MHuQUitUNo*XeimH
		ieavpFcl@gq!tA@9
		
	2. strong
	
		Lul]M\<
		rW3xW|&
		i<qbtHf
		
		HX9hr5eO
		;4&wvkj)
		
		mp.G*6"ZL\\PFo]|z
		j{RzI]q~G>Q6i^s^=:
		.(Gt^R*wR?k^8`o}+&
		
- Passphrases:

	1. humanized
	
		WiCHad eQeq1$
		M8aVK@ eja*a

		*5a e1mHI NUdUm
		R16 y3pee Dz0Od
		
	2. strong
	
		mM47d| q{5Ls
		_+Of!% &V1c&

## Version

v0.0.2-alpha

## Install

##### As application

Get current build at [release page](https://github.com/hIMEI29A/pwgator/releases).

Or run

	wget https://github.com/hIMEI29A/pwgator/releases/download/\
	v0.0.2-alpha/pwgator-0.0.2-alpha-amd64.deb && sudo dpkg -i pwgator-0.0.2-alpha-amd64.deb
		

##### As package

```sh
go get github.com/hIMEI29A/pwgator/pwgator
```

## Usage

Call **pwgator** with `-h` or `--help` option to get full usage instruction.

	$ pwgator -h

	pwgator - humanized passphrases generator.

	Usage:
	  pwgator [-h | --help] | [--version]
	  pwgator [LENGTH] [-s | --strong] [-p | --phrase]

	Arguments:
	  LENGTH               Secret's length (OPTIONAL). Default is 8 for word, 2 for phrase.

	Options:
	  -h --help            Show this screen.
	  --version            Show version.
	  -s --strong          Non-humanized generation.
	  -p --phrase          Generate passphrases.

Generation of default humanized passwords:

	pwgator
	
Generation of default humanized passphrases:

	pwgator -p
	
Passwords with length = 10:

	pwgator 10
	
Length = 10, strong generation:

	pwgator 10 -s
	
Passphrases with 4 words length and in strong mode:

	pwgator 4 -p -s
	
To use **pwgen** as packege in your applications, write something like this:

```go
package main

import (
	"fmt"
	
	"github.com/hIMEI29A/pwgator/pwgator"
)	
	func main() {
		length := 8
		
		// Means false
		phrases := 0
		
		strong := true
		
		// Generate new pwgator with phrases boolean option, 
		//int length and boolean strong option
		app := pwgator.NewApp(phrases, length, strong)
		app.Generate()

		fmt.Println(app.String())
	}
```