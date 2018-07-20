package im

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/wst-libs/wstsdk/errors"
	"github.com/wst-libs/wstsdk/im/sdk"
	"github.com/wst-libs/wstsdk/utils"
)

func SetOutPutHeader(ctx *context.Context) {
	ctx.Output.Header("Connection", "close")
	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Header("Server", "wst-im:v1.0")
}

func ParseInputHander(ctx *context.Context) error {
	if !strings.Contains(ctx.Input.Header("Content-Type"), "application/json") {
		log.Println("not application/json.")
		// return error.Error()("Content-Type not application/json")
		// return errors.New("Content-Type is application/json")
	}
	return nil
}

///
/// Processing user registration
///
func RegisterUsersHandler(ctx *context.Context) []byte {
	var request RequestRegisteredUsers

	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	uid := ctx.Input.Param(":uid")
	s := sdk.NewRCServer()
	token := s.GetTokenFromUser(uid, request.RegisteredUsers.Name, request.RegisteredUsers.Portrait)

	v := ResponseRegisteredUsers{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  request.SeqNum,
			From:    request.From,
			To:      request.To,
			// Type:    request.Type,
			Number: request.Number,
			Code:   0,
		},
		utils.TOKEN{
			Token: token,
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

///
/// Processing create session
///
func CreateSessionHandler(ctx *context.Context) []byte {
	var request RequestCreateSession

	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	sid := ctx.Input.Param(":sid")
	s := sdk.NewRCServer()
	s.CreateChatRoom(sid, request.Name)

	v := ResponseCreateSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
		utils.ID{
			Id: sid,
		},
	}
	out, err := json.Marshal(v)
	if err != nil {
	}
	return out
}

///
/// Processing delete session
///
func DeleteSessionHandler(ctx *context.Context) []byte {
	var request RequestDelSession
	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}
	ids := make([]string, request.Size)
	for i, v := range request.List {
		ids[i] = v.Id
	}
	s := sdk.NewRCServer()
	s.DeleteChatRoom(ids)

	v := ResponseDelSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

///
/// Processing delete session by session id
///
func DeleteSessionByIDHandler(ctx *context.Context) []byte {
	var request RequestDelSession
	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	sid := ctx.Input.Param(":sid")

	s := sdk.NewRCServer()
	s.DeleteChatRoom([]string{sid})

	v := ResponseDelSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

///
/// Processing delete session by user id
///
func DeleteSessionByUIDHandler(ctx *context.Context) []byte {
	var request RequestDelSession
	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	// s := sdk.NewRCServer()
	v := ResponseDelSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

///
/// Processing get session
///
func GetSessionHandler(ctx *context.Context) []byte {

	v := ResponseGetSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
		GetSession{},
	}
	// s := sdk.NewRCServer()
	// s.GetAllChatRoom()
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

///
/// Processing get session infomation by session id
///
func GetSessionByIDHandler(ctx *context.Context) []byte {

	sid := ctx.Input.Param(":sid")
	s := sdk.NewRCServer()
	s.GetChatRoomById([]string{sid})

	v := ResponseGetSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
		GetSession{},
	}

	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

///
/// Processing get session by user id
///
func GetSessionByUIDHandler(ctx *context.Context) []byte {
	v := ResponseGetSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
		GetSession{},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

///
/// Processing get users by session id
///
func GetUsersBySessionIDHandler(ctx *context.Context) []byte {
	v := ResponseGetSessionUsers{
		utils.ResponseCommon{
			Version: "V1.0",
		},
		GetSessionUsers{},
	}

	uid := ctx.Input.Param(":uid")
	s := sdk.NewRCServer()
	s.GetUsersByRoomId(uid)

	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

func PutSessionByUIDHandler(ctx *context.Context) []byte {
	var request RequestJoinSession
	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	uid := ctx.Input.Param(":uid")
	sid := ctx.Input.Param(":sid")
	s := sdk.NewRCServer()
	s.JoinRoomByUserId(uid, sid)

	v := ResponseJoinSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

func PostMessageToUserByIDHandler(ctx *context.Context) []byte {
	var request RequestSendMessage
	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	sourceId := request.Uid
	targetId := ctx.Input.Param(":uid")
	content := request.Content

	log.Println("source id: ", sourceId)
	log.Println("target id: ", targetId)
	log.Println("content is: ", content)

	s := sdk.NewRCServer()
	s.SendMsgUserToUsers(sourceId, targetId, content)

	v := ResponseSendMessage{
		utils.ResponseCommon{
			Version: "V1.0",
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

func PostMessageToSessionByIDHandler(ctx *context.Context) []byte {
	var request RequestSendMessage
	err := json.Unmarshal(ctx.Input.RequestBody, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	v := ResponseSendMessage{
		utils.ResponseCommon{
			Version: "V1.0",
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}
