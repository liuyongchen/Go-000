package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// cfg := web.

	_ = []*MainController{
		&MainController{"1", "2"},
	}

	router := web.NewHttpSever()
	router.Router("/", nil, "get:Get")
	router.Run(":8080")
}

type MainController struct {
	// web.Controller
	id   string
	name string
}

// address: http://localhost:8080 GET
// func (ctrl *MainController) Get() {
//
// 	// beego-example/views/hello_world.html
// 	ctrl.TplName = "hello_world.html"
// 	ctrl.Data["name"] = "Get()"
//
// 	// don't forget this
// 	_ = ctrl.Render()
// }
