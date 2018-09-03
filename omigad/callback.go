package omigad

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	"github.com/wst-libs/wst-sdk/sdk/manager"
)

type CallBack struct {
}

type ReqBody struct {
	Id string `json:"id"`
}
type ResSuccess struct {
	Code int64 `json:"code"`
}
type ResFailed struct {
	Code int64 `json:"code"`
}

func ResponseSuccess() []byte {
	out := ResSuccess{
		Code: 0,
	}
	res, _ := json.Marshal(out)
	return res
}
func ResponseFailed() []byte {
	out := ResFailed{
		Code: 404,
	}
	res, _ := json.Marshal(out)
	return res
}

func CallBackHandler(ctx *context.Context) []byte {
	log.Println("callback: ", ctx.Input.RequestBody)
	var req ReqBody
	if json.Unmarshal(ctx.Input.RequestBody, req) != nil {
		log.Println("Failed to json unmarshal.")
		return ResponseFailed()
	}

	manager.Update(beego.AppConfig.String("managerurl")+"/"+req.Id, 3)

	return ResponseSuccess()
}
