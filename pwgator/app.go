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

const (
	MAX_GOROUTINES = 100
	BUFFSIZE       = 200
)

// Constants for password or passphrase generation
const (
	PASSWORD APP_T = iota
	PASSPHRASE
)

type APP_T int

// App is a type for options generation.
type App struct {
	// Generated is a field for words or phrases produced
	Generated []string
	// AppType is a const for password or passphrase
	AppType APP_T
	// IntValis a integer length in symbols (for word) or words(for phrases)
	IntVal int
	// Boolean option for strong/humanized generation
	Strong bool
	// Boolean option for random/templated passphrase
	Random bool
}

// NewApp instantiates App datatype
func NewApp(t, val int, s, r bool) *App {
	app := &App{}
	app.AppType = APP_T(t)
	app.IntVal = val
	app.Strong = s
	app.Random = r

	return app
}

// Generate generates 100 strings
func (a *App) Generate() {
	var generated []string
	var chanString = make(chan string, BUFFSIZE)

	for i := 0; i < MAX_GOROUTINES; i++ {
		go func(ch chan string) {
			var s string
			secret := NewSecret()

			if a.AppType == PASSWORD {
				s = secret.PassWord(a.IntVal, a.Strong)
			} else {
				s = secret.PassPhrase(a.IntVal, a.Strong, a.Random)
			}

			ch <- s
		}(chanString)
	}

	for len(generated) < MAX_GOROUTINES {
		select {
		case recvd := <-chanString:
			generated = append(generated, recvd)

		default:
		}
	}

	a.Generated = generated
}

// String turns generated []string to string
// Each new line is ended with "\n"
func (a *App) String() string {
	var s string

	for i := range a.Generated {
		s = s + a.Generated[i] + "\n"
	}

	return s
}
