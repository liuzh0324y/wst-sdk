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
func CreateSessionHandler(body []byte) []byte {
	var request RequestCreateSession

	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	v := ResponseCreateSession{
		utils.ResponseCommon{
			Version: "V1.0",
		},
		utils.ID{
			Id: "",
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
func DeleteSessionHandler(body []byte) []byte {
	var request RequestDelSession
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

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
func DeleteSessionByIDHandler(body []byte) []byte {
	var request RequestDelSession
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

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
func DeleteSessionByUIDHandler(body []byte) []byte {
	var request RequestDelSession
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

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
func GetSessionHandler(body []byte) []byte {

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
/// Processing get session by session id
///
func GetSessionByIDHandler(body []byte) []byte {
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
func GetSessionByUIDHandler(body []byte) []byte {
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
func GetUsersBySessionIDHandler(body []byte) []byte {
	v := ResponseGetSessionUsers{
		utils.ResponseCommon{
			Version: "V1.0",
		},
		GetSessionUsers{},
	}

	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

func PutSessionByUIDHandler(body []byte) []byte {
	var request RequestJoinSession
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

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

func PostMessageToUserByIDHandler(body []byte) []byte {
	var request RequestSendMessage
	err := json.Unmarshal(body, &request)
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

func PostMessageToSessionByIDHandler(body []byte) []byte {
	var request RequestSendMessage
	err := json.Unmarshal(body, &request)
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
