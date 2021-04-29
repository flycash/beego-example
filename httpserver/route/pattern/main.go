// Copyright 2020 
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) HelloWorld()  {
	u.Ctx.WriteString("Hello, world")
}

func (u *UserController) HelloWorld1()  {
	u.Ctx.WriteString("Hello, mid")
}

func main() {
	web.RouterGet("api/*/mid/name", (*UserController).HelloWorld1)
	web.RouterGet("api/*/name", (*UserController).HelloWorld)

	web.Run()
}
