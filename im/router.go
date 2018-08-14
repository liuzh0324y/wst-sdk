package im

import (
	"github.com/astaxie/beego"
)

func init() {
	// Mapping user routing
	// deprecated
	// beego.Router("/api/v1/im/user/:uid/register", &Controller{}, "post:RegisterUsers")
	// beego.Router("/api/v1/im/user/:uid", &Controller{}, "get:GetSessionByUID")
	beego.Router("/api/v1/im/user/:uid/sendmessage", &Controller{}, "post:PostMessageToUserByID")
	beego.Router("/api/v1/im/register", &Controller{}, "post:RegisterUsers")

	// Mapping session routing
	beego.Router("/api/v1/im/session", &Controller{}, "delete:DeleteSession")
	beego.Router("/api/v1/im/session/:sid", &Controller{}, "put:CreateSessionById;delete:DeleteSessionByID")

	// Mapping message routing
	beego.Router("/api/v1/im/message", &Controller{}, "put:ReqMsgCtl")

	// query
	// deprecated
	// beego.Router("/api/v1/im/session", &Controller{}, "get:GetAllSession")
	// beego.Router("/api/v1/im/session/:sid", &Controller{}, "get:GetSessionByID")
	// beego.Router("/api/v1/im/session/:sid/user", &Controller{}, "get:GetUsersBySessionID")
	// beego.Router("/api/v1/im/session/:sid/user/:uid", &Controller{}, "put:PutSessionByUID")
	// beego.Router("/api/v1/im/session/:sid/user/:uid", &Controller{}, "delete:DeleteSessionByUID")
	// beego.Router("/api/v1/im/session/:sid/sendmessage", &Controller{}, "post:PostMessageToSessionByID")
}
