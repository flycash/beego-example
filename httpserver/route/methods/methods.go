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

func (u *UserController) GetUserById() {
	u.Ctx.WriteString("GetUserById")
}

func (u *UserController) UpdateUser() {
	u.Ctx.WriteString("UpdateUser")
}

func (u *UserController) UserHome() {
	u.Ctx.WriteString("UserHome")
}

func (u *UserController) DeleteUser() {
	u.Ctx.WriteString("DeleteUser")
}

func (u *UserController) HeadUser() {
	u.Ctx.WriteString("HeadUser")
}

func (u *UserController) OptionUsers() {
	u.Ctx.WriteString("OptionUsers")
}

func (u *UserController) PatchUsers() {
	u.Ctx.WriteString("PatchUsers")
}

func (u *UserController) PutUsers() {
	u.Ctx.WriteString("PutUsers")
}

func main() {

	// get http://localhost:8080/api/user/123
	web.RouterGet("api/user/:id", (*UserController).GetUserById)

	// post http://localhost:8080/api/user/update
	web.RouterPost("api/user/update", (*UserController).UpdateUser)

	// http://localhost:8080/api/user/home
	web.RouterAny("api/user/home", (*UserController).UserHome)

	// delete http://localhost:8080/api/user/delete
	web.RouterDelete("api/user/delete", (*UserController).DeleteUser)

	// head http://localhost:8080/api/user/head
	web.RouterHead("api/user/head", (*UserController).HeadUser)

	// patch http://localhost:8080/api/user/options
	web.RouterOptions("api/user/options", (*UserController).OptionUsers)

	// patch http://localhost:8080/api/user/patch
	web.RouterPatch("api/user/patch", (*UserController).PatchUsers)

	// put http://localhost:8080/api/user/put
	web.RouterPut("api/user/put", (*UserController).PutUsers)

	web.Run()
}
