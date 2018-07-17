package im

import (
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

///
/// Register users
///
func (this *Controller) RegisterUsers() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(RegisterUsersHandler(this.Ctx.Input.RequestBody))
}

///
/// Create session
///
func (this *Controller) CreateSession() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(CreateSessionHandler(this.Ctx.Input.RequestBody))
}

///
/// Delete all session
///
func (this *Controller) DeleteSession() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(DeleteSessionHandler(this.Ctx.Input.RequestBody))
}

///
/// Delete session by session id
///
func (this *Controller) DeleteSessionByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(DeleteSessionByIDHandler(this.Ctx.Input.RequestBody))
}

///
/// Delete session by user id
///
func (this *Controller) DeleteSessionByUID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(DeleteSessionByUIDHandler(this.Ctx.Input.RequestBody))
}

///
/// Get all session
///
func (this *Controller) GetSession() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetSessionHandler(this.Ctx.Input.RequestBody))
}

///
/// Get session by session id
///
func (this *Controller) GetSessionByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetSessionByIDHandler(this.Ctx.Input.RequestBody))
}

///
/// Get session by user id
///
func (this *Controller) GetSessionByUID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetSessionByUIDHandler(this.Ctx.Input.RequestBody))
}

///
/// Get users by session id
///
func (this *Controller) GetUsersBySessionID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetUsersBySessionIDHandler(this.Ctx.Input.RequestBody))
}

///
/// Put session by user id
///
func (this *Controller) PutSessionByUID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(PutSessionByUIDHandler(this.Ctx.Input.RequestBody))
}

///
/// Post message to user by session id
///
func (this *Controller) PostMessageToUserByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(PostMessageToUserByIDHandler(this.Ctx.Input.RequestBody))
}

///
/// Post message to session by session id
///
func (this *Controller) PostMessageToSessionByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(PostMessageToSessionByIDHandler(this.Ctx.Input.RequestBody))
}
