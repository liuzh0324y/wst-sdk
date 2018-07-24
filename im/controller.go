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
	this.Ctx.Output.Body(RegisterUsersHandler(this.Ctx))
}

///
/// Create session
///
func (this *Controller) CreateSessionById() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(CreateSessionHandler(this.Ctx))
}

///
/// Delete all session
///
func (this *Controller) DeleteSession() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(DeleteSessionHandler(this.Ctx))
}

///
/// Delete session by session id
///
func (this *Controller) DeleteSessionByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(DeleteSessionByIDHandler(this.Ctx))
}

///
/// Delete session by user id
///
func (this *Controller) DeleteSessionByUID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(DeleteSessionByUIDHandler(this.Ctx))
}

///
/// Get all session
///
func (this *Controller) GetAllSession() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetSessionHandler(this.Ctx))
}

///
/// Get session by session id
///
func (this *Controller) GetSessionByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetSessionByIDHandler(this.Ctx))
}

///
/// Get session by user id
///
func (this *Controller) GetSessionByUID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetSessionByUIDHandler(this.Ctx))
}

///
/// Get users by session id
///
func (this *Controller) GetUsersBySessionID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(GetUsersBySessionIDHandler(this.Ctx))
}

///
/// Put session by user id
///
func (this *Controller) PutSessionByUID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(PutSessionByUIDHandler(this.Ctx))
}

///
/// Post message to user by session id
///
func (this *Controller) PostMessageToUserByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(PostMessageToUserByIDHandler(this.Ctx))
}

///
/// Post message to session by session id
///
func (this *Controller) PostMessageToSessionByID() {
	SetOutPutHeader(this.Ctx)
	this.Ctx.Output.Body(PostMessageToSessionByIDHandler(this.Ctx))
}
