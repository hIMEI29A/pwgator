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

import (
	"fmt"
)

const MAX_GOROUTINES = 100

const (
	PASSWORD APP_T = iota
	PASSPHRASE
)

type APP_T int

type App struct {
	Generated []string
	AppType   APP_T
	IntVal    int
}

func NewApp(t, val int) *App {
	app := &App{}
	app.AppType = APP_T(t)
	app.IntVal = val

	return app
}

func (a *App) Generate() []string {

}
