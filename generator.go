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
	maxGoroutines = 100
	buffsize      = 200
)

// Const represented type of secret: password or passphrase
const (
	Password GeneratorType = iota
	Passphrase
)

// GeneratorType is a type of secret: password or passphrase
type GeneratorType int

// Generator is a main type for passwords and passphrases generation
type Generator struct {
	// Generated is a field for words or phrases produced
	Generated []string
	// GenType may be password or passphrase
	GenType GeneratorType
	// IntVal is a secrets length in symbols (for word) or in words(for phrases)
	IntVal int
	// Strong is a boolean option for strong/humanized generation
	Strong bool
	// Random is a boolean option for random/templated passphrase
	Random bool
}

// NewGenerator instantiates Generator
func NewGenerator(t, val int, s, r bool) *Generator {
	return &Generator{
		GenType: GeneratorType(t),
		IntVal:  val,
		Strong:  s,
		Random:  r,
	}
}

// generate generates num of passwowrds or passpgrases
func (g *Generator) generate(num int) {
	var generated []string
	var chanString = make(chan string, buffsize)

	for i := 0; i < num; i++ {
		go func(ch chan string) {
			var s string
			sec := newSecret()

			if g.GenType == Password {
				s = sec.passWord(g.IntVal, g.Strong)
			} else {
				s = sec.passPhrase(g.IntVal, g.Strong, g.Random)
			}

			ch <- s
		}(chanString)
	}

	for len(generated) < num {
		select {
		case recvd := <-chanString:
			generated = append(generated, recvd)

		default:
		}
	}

	g.Generated = generated
}

// Generate generates 100 passwowrds or passpgrases
func (g *Generator) Generate() {
	g.generate(maxGoroutines)
}

func (g *Generator) GenerateByValue(value int) {
	g.generate(value)
}

// GetGenerated returns slice of generated strings
func (g *Generator) GetGenerated() []string {
	return g.Generated
}

// GetString returns Generator.Generated as string
func (g *Generator) GetString() string {
	var s string

	for i := range g.Generated {
		s = s + g.Generated[i] + "\n"
	}

	return s
}
