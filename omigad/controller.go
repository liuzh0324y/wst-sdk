package omigad

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type Controller struct {
	beego.Controller
}

func SetOutPutHeaders(ctx *context.Context) {
	ctx.Output.Header("Connection", "close")
	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Header("Server", "omigad:V1.0")
}
func (t *Controller) PutFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(PutFileHandler(t.Ctx))
}

func (t *Controller) GetFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(GetFileHandler(t.Ctx))
}

func (t *Controller) UpdateFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(UpdateFileHandler(t.Ctx))
}

func (t *Controller) DeleteFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(DeleteFileHandler(t.Ctx))
}

func (t *Controller) CallBack() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(CallBackHandler(t.Ctx))
}

func (t *Controller) GetUrlForFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(GetUrlForFileHandler(t.Ctx))
}
