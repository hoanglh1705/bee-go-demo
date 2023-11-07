package main

import (
	"bee-go-demo/config"
	filecontroller "bee-go-demo/httpserver/controller/file"
	homecontroller "bee-go-demo/httpserver/controller/home"
	usercontroller "bee-go-demo/httpserver/controller/user"
	"fmt"

	mydb "bee-go-demo/db"

	apmbeego "github.com/opentracing-contrib/beego"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Load config err", err)
		panic(err)
	}

	web.BConfig.CopyRequestBody = true
	mydb.Init(cfg)
	// Init Controllers
	userCtrl := usercontroller.NewUserController("/users")
	fileCtrl := filecontroller.NewFileController("/files")
	mainCtrl := homecontroller.NewHomeController("/")

	// Register Controllers
	web.Router(userCtrl.GetPath(), userCtrl)
	web.Router(fileCtrl.GetPath(), fileCtrl)
	web.Router(mainCtrl.GetPath(), mainCtrl)

	web.RunWithMiddleWares("localhost:8080", apmbeego.Middleware("bee-go-demo"))
	// lambda.Start()

}
