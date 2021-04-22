package main

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"

)

func main() {
	web.Router("/", &MainController{})
	web.Router("/redirect", &RedirectController{})
	web.Run()
}

type RedirectController struct {
	web.Controller
}

func (r *RedirectController) Post() {
	// redirect to /, must use get method
	r.Redirect("/", http.StatusSeeOther)
}

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString("hello world")
}
