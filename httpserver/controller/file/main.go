package file

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type FileController struct {
	web.Controller
	path string
}

func NewFileController(path string) *FileController {
	ctrl := &FileController{}
	ctrl.path = path

	web.Router(ctrl.path+"/upload", ctrl, "post:Upload")
	web.Router(ctrl.path+"/upload", ctrl, "get:UploadPage")
	web.Router(ctrl.path+"/save", ctrl, "post:Save")

	return ctrl
}

func (c *FileController) GetPath() (path string) {
	return c.path
}

// GET http://localhost:8080/upload
// and you will see the upload page
func (ctrl *FileController) UploadPage() {
	ctrl.TplName = "upload.html"
}

// POST http://localhost:8080/upload
// you will see "success"
// and the file name (actual file name, not the key you use in GetFile
func (ctrl *FileController) Upload() {
	// key is the file name
	file, fileHeader, err := ctrl.GetFile("upload.txt")
	if err != nil {
		logs.Error("save file failed, ", err)
		ctrl.Ctx.Output.Body([]byte(err.Error()))
	} else {
		// don't forget to close
		defer file.Close()

		logs.Info(fileHeader.Filename)
		ctrl.Ctx.Output.Body([]byte("success"))
	}

}

// POST http://localhost:8080/save
// you will see the file /tmp/upload.txt and "success"
// and if you run this on Windows platform, don't forget to change the target file path
func (ctrl *FileController) Save() {
	err := ctrl.SaveToFile("save.txt", "/tmp/upload.txt")
	if err != nil {
		logs.Error("save file failed, ", err)
		ctrl.Ctx.Output.Body([]byte(err.Error()))
	} else {
		ctrl.Ctx.Output.Body([]byte("success"))
	}
}
