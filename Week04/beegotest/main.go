package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main()  {
	cfg := web.


	ctrl := &MainController{}

	router := web.NewHttpSever()
	router.Router("/", ctrl, "get:Get")
	router.Run(":8080")
}

type MainController struct {
	web.Controller
}

// address: http://localhost:8080 GET
func (ctrl *MainController) Get() {

	// beego-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "Get()"

	// don't forget this
	_ = ctrl.Render()
}