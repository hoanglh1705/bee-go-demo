package main

import (
	"bee-go-demo/config"
	filecontroller "bee-go-demo/httpserver/controller/file"
	homecontroller "bee-go-demo/httpserver/controller/home"
	usercontroller "bee-go-demo/httpserver/controller/user"
	userrepository "bee-go-demo/httpserver/repository/user"
	"fmt"
	"net/http"
	"text/template"

	mydb "bee-go-demo/db"

	apmbeego "github.com/opentracing-contrib/beego"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/prometheus"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Load config err", err)
		panic(err)
	}

	web.BConfig.CopyRequestBody = true
	mydb.Init(cfg)
	o := orm.NewOrm()
	userRepository := userrepository.NewUserRepository(o)

	// Init Controllers
	userCtrl := usercontroller.NewUserController("/users", userRepository)
	fileCtrl := filecontroller.NewFileController("/files")
	mainCtrl := homecontroller.NewHomeController("/")

	// Register Controllers
	web.ErrorHandler("404", page_not_found)
	web.Router(userCtrl.GetPath(), userCtrl)
	web.Router(fileCtrl.GetPath(), fileCtrl)
	web.Router(mainCtrl.GetPath(), mainCtrl)

	fb := &prometheus.FilterChainBuilder{}
	web.InsertFilterChain("/*", fb.FilterChain)

	web.RunWithMiddleWares(fmt.Sprintf("%v:%d", cfg.Host, cfg.Port), apmbeego.Middleware("bee-go-demo"))
	// lambda.Start()

}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(web.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
