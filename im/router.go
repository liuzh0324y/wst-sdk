package im

import (
	"github.com/astaxie/beego"
)

func init() {
	// Mapping user routing
	beego.Router("/api/v1/im/user/:uid/register", &Controller{}, "post:RegisterUsers")
	beego.Router("/api/v1/im/user/:uid", &Controller{}, "get:GetSessionByUID")
	beego.Router("/api/v1/im/user/:uid/sendmessage", &Controller{}, "post:PostMessageToUserByID")

	// Mapping session routing
	// create
	beego.Router("/api/v1/im/session/:sid", &Controller{}, "put:CreateSessionById")
	// delete
	beego.Router("/api/v1/im/session/:sid", &Controller{}, "delete:DeleteSessionByID")
	beego.Router("/api/v1/im/session", &Controller{}, "delete:DeleteSession")
	// query
	beego.Router("/api/v1/im/session", &Controller{}, "get:GetAllSession")
	beego.Router("/api/v1/im/session/:sid", &Controller{}, "get:GetSessionByID")
	beego.Router("/api/v1/im/session/:sid/user", &Controller{}, "get:GetUsersBySessionID")
	beego.Router("/api/v1/im/session/:sid/user/:uid", &Controller{}, "put:PutSessionByUID")
	beego.Router("/api/v1/im/session/:sid/user/:uid", &Controller{}, "delete:DeleteSessionByUID")
	beego.Router("/api/v1/im/session/:sid/sendmessage", &Controller{}, "post:PostMessageToSessionByID")
}
