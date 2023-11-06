package user

import (
	"bee-go-demo/httpserver/controller/dto"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
	path string
}

func NewUserController(path string) *UserController {
	ctrl := &UserController{}
	ctrl.path = path
	web.Router(ctrl.path+"/:name", ctrl, "get:Query")
	return ctrl
}

func (c *UserController) GetPath() (path string) {
	return c.path
}

func (u *UserController) HelloWorld() {
	u.Ctx.WriteString("Hello world!")
}

// address: http://localhost:8080/users Post
func (ctrl *UserController) Post() {
	input := dto.CreateUserRequestDTO{}

	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &input); err != nil {
		ctrl.Data["json"] = err.Error()
	}

	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}

// address: http://localhost:8080/users Get
func (ctrl *UserController) Get() {
	input := dto.UserResponseDTO{
		Name:  "admin",
		Email: "admin@bee-go-demo.com",
	}

	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}

// Get
// address: http://localhost:8080/users/:name GET
// Get use default value
func (ctrl *UserController) Query() {
	name := ctrl.GetString("name", "Tom")
	ctrl.Ctx.WriteString("Hello " + name)
}
