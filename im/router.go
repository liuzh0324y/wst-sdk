package im

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/im/user/:uid/register", &Controller{}, "post:RegisterUsers")
	beego.Router("/api/v1/im/session/:sid", &Controller{}, "get:GetSessionByID;put:CreateSession;delete:DeleteSessionByID")
	beego.Router("/api/v1/im/session", &Controller{}, "get:GetSession;delete:DeleteSession")
	beego.Router("/api/v1/im/user/:uid", &Controller{}, "get:GetSessionByUID")
	beego.Router("/api/v1/im/session/:sid/user", &Controller{}, "get:GetUsersBySessionID")
	beego.Router("/api/v1/im/session/:sid/user/:uid", &Controller{}, "put:PutSessionByUID;delete:DeleteSessionByUID")
	beego.Router("/api/v1/im/user/:uid/sendmessage", &Controller{}, "post:PostMessageToUserByID")
	beego.Router("/api/v1/im/session/:sid/sendmessage", &Controller{}, "post:PostMessageToSessionByID")
}
