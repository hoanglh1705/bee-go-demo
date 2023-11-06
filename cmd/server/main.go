package main

import (
	filecontroller "bee-go-demo/httpserver/controller/file"
	homecontroller "bee-go-demo/httpserver/controller/home"
	usercontroller "bee-go-demo/httpserver/controller/user"

	mydb "bee-go-demo/db"

	apmbeego "github.com/opentracing-contrib/beego"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.BConfig.CopyRequestBody = true
	mydb.Init()
	// Init Controllers
	userCtrl := usercontroller.NewUserController("/users")
	fileCtrl := filecontroller.NewFileController("/files")
	mainCtrl := homecontroller.NewHomeController("/")

	// Register Controllers
	web.Router(userCtrl.GetPath(), userCtrl)
	web.Router(fileCtrl.GetPath(), fileCtrl)
	web.Router("/", mainCtrl)

	web.RunWithMiddleWares("localhost:8080", apmbeego.Middleware("bee-go-demo"))
	// lambda.Start()

}
