package home

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/prometheus"
)

type HomeController struct {
	web.Controller
	path string
}

func (c *HomeController) GetPath() (path string) {
	return c.path
}

func NewHomeController(path string) *HomeController {
	// we start admin service
	// Prometheus will fetch metrics data from admin service's port
	web.BConfig.Listen.EnableAdmin = true

	web.BConfig.AppName = "my app"

	ctrl := &HomeController{}
	ctrl.path = path

	web.Router("/hello", ctrl, "get:Hello")
	fb := &prometheus.FilterChainBuilder{}
	web.InsertFilterChain("/*", fb.FilterChain)

	return ctrl
}

func (ctrl *HomeController) Hello() {
	ctrl.Ctx.ResponseWriter.Write([]byte("Hello, world"))
}
