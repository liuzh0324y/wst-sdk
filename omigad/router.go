package omigad

import (
	"log"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/cloudstorage/file", &Controller{}, "put:PutFile;get:GetFile;post:UpdateFile;delete:DeleteFile")
	beego.Router("/api/v1/cloudstorage/callback", &Controller{}, "post:CallBack")
	beego.Router("/api/v1/cloudstorage/uploadinfo", &Controller{}, "get:GetUrlForFile")
	log.Println("Initialize Router.")
}
