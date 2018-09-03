package im

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wst-libs/wst-sdk/errors"
	"github.com/wst-libs/wst-sdk/im/sdk"
	"github.com/wst-libs/wst-sdk/utils"
)

// Run is a start mothod
func Run() {
	var addr string
	err := getconf()
	if err != nil {
		addr = ":18010"
	} else {
		addr = ":" + config.Server.Httpport
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/im/register", register)
	router.HandleFunc("/api/v1/im/session", session)
	router.HandleFunc("/api/v1/im/session/{sid}", sessionbyid)
	router.HandleFunc("/api/v1/im/message", sendmsg)

	log.Fatal(http.ListenAndServe(addr, router))
}

// register user
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Connection", "close")
		w.Write(errors.NotSupportMethod())
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Connection", "close")
	w.Write(registerHandler(body))
	defer r.Body.Close()
}

// delete: delete session
func session(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Connection", "close")
		w.Write(errors.NotSupportMethod())
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Connection", "close")
	w.Write(delAllSessionHandler(body))
	defer r.Body.Close()
}

// put: create session by sid; delete: delete session by sid
func sessionbyid(w http.ResponseWriter, r *http.Request) {
	var outBody []byte
	vars := mux.Vars(r)
	sid := vars["sid"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if r.Method == "PUT" {
		outBody = createSessionByID(sid, body)
	} else if r.Method == "DELETE" {
		outBody = delSessionByID(sid, body)
	} else {
		outBody = errors.NotSupportMethod()
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Connection", "close")
	w.Write(outBody)
}

// send message to target
func sendmsg(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Connection", "close")
	w.Write(sendmsgHandler(body))
}

// Processing user registration
func registerHandler(body []byte) []byte {
	// Parse request body to json
	var request RequestRegisteredUsers
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	// Check request common param
	com := errors.IsCommonErr(request.RequestCommon)
	if com.Code != 0 {
		outerr, _ := json.Marshal(com)
		return outerr
	}

	// Parse request uid
	// uid := ctx.Input.Param(":uid")
	// err = errors.VerifyUid(uid)
	// if err != nil {
	// 	ret := utils.ResponseCommon{
	// 		Code:    errors.UidErr,
	// 		Message: err.Error(),
	// 		Version: request.RequestCommon.Version,
	// 		SeqNum:  request.RequestCommon.SeqNum,
	// 		From:    request.RequestCommon.From,
	// 		To:      request.RequestCommon.To,
	// 		Type:    request.RequestCommon.Type,
	// 		Number:  request.RequestCommon.Number,
	// 	}
	// 	out, _ := json.Marshal(ret)
	// 	return out
	// }
	_, err = checkRegisterCommon(request.RegisteredUsers)
	if err != nil {
		return errors.ImplementErr(errors.UserInfoErr, request.RequestCommon, err.Error())
	}

	s := sdk.NewRCServer()
	token, err := s.GetTokenFromUser(request.RegisteredUsers.Id, request.RegisteredUsers.Name, request.RegisteredUsers.Portrait)
	if err != nil {
		return errors.ImplementErr(errors.UserTokenErr, request.RequestCommon, err.Error())
	}

	v := ResponseRegisteredUsers{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  request.RequestCommon.SeqNum,
			From:    request.RequestCommon.From,
			To:      request.RequestCommon.To,
			Type:    request.RequestCommon.Type,
			Number:  request.RequestCommon.Number,

			Code:    0,
			Message: "success",
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

func delAllSessionHandler(body []byte) []byte {
	// Parse request body to json
	var request RequestDelSession
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}
	// Check request common params
	com := errors.IsCommonErr(request.RequestCommon)
	if com.Code != 0 {
		outerr, _ := json.Marshal(com)
		return outerr
	}
	// Get session id by request params
	ids := make([]string, request.Size)
	for i, v := range request.List {
		ids[i] = v.Id
	}
	s := sdk.NewRCServer()
	err = s.DeleteChatRoom(ids)
	if err != nil {
		return errors.ImplementErr(errors.DeleteSessionErr, request.RequestCommon, err.Error())
	}

	v := ResponseDelSession{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  request.RequestCommon.SeqNum,
			From:    request.RequestCommon.From,
			To:      request.RequestCommon.To,
			Type:    request.RequestCommon.Type,
			Number:  request.RequestCommon.Number,

			Message: "success",
			Code:    0,
		},
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

func createSessionByID(sid string, body []byte) []byte {
	// Parse request body to json
	var request RequestCreateSession
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	// Check request common param
	com := errors.IsCommonErr(request.RequestCommon)
	if com.Code != 0 {
		outerr, _ := json.Marshal(com)
		return outerr
	}

	// Get session id by uri
	err = errors.VerifyUid(sid)
	if err != nil {
		ret := utils.ResponseCommon{
			Code:    errors.UidErr,
			Message: err.Error(),
			Version: request.RequestCommon.Version,
			SeqNum:  request.RequestCommon.SeqNum,
			From:    request.RequestCommon.From,
			To:      request.RequestCommon.To,
			Type:    request.RequestCommon.Type,
			Number:  request.RequestCommon.Number,
		}
		out, _ := json.Marshal(ret)
		return out
	}

	s := sdk.NewRCServer()
	err = s.CreateChatRoom(sid, request.Name)
	if err != nil {
		return errors.ImplementErr(errors.CreateSessionErr, request.RequestCommon, err.Error())
	}

	v := ResponseCreateSession{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  request.RequestCommon.SeqNum,
			From:    request.RequestCommon.From,
			To:      request.RequestCommon.To,
			Type:    request.RequestCommon.Type,
			Number:  request.RequestCommon.Number,

			Message: "success",
			Code:    0,
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

func delSessionByID(sid string, body []byte) []byte {
	// Parse request body to json
	var request RequestDelSession
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}
	// Check request common
	com := errors.IsCommonErr(request.RequestCommon)
	if com.Code != 0 {
		outerr, _ := json.Marshal(com)
		return outerr
	}

	s := sdk.NewRCServer()
	err = s.DeleteChatRoom([]string{sid})
	if err != nil {
		return errors.ImplementErr(errors.DeleteSessionErr, request.RequestCommon, err.Error())
	}

	v := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  request.RequestCommon.SeqNum,
		From:    request.RequestCommon.From,
		To:      request.RequestCommon.To,
		Type:    request.RequestCommon.Type,
		Number:  request.RequestCommon.Number,

		Message: "success",
		Code:    0,
	}
	out, err := json.Marshal(v)
	if err != nil {

	}
	return out
}

func sendmsgHandler(body []byte) []byte {
	var req ReqMsg
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error: ", err.Error())
		return errors.ParseJsonFailed()
	}

	com := errors.IsCommonErr(req.RequestCommon)
	if com.Code != 0 {
		outerr, _ := json.Marshal(com)
		return outerr
	}

	s := sdk.NewRCServer()
	switch req.Data.Type {
	case "private":
		s.SendMsgPrivate(req.Data.SourceId, req.Data.TargetId, req.Data.Content)
	case "group":
		s.SendMsgGroup(req.Data.SourceId, req.Data.TargetId, req.Data.Content)
	case "chatroom":
		s.SendMsgRoom(req.Data.SourceId, req.Data.TargetId, req.Data.Content)
	case "system":
		s.SendMsgSystem(req.Data.SourceId, req.Data.TargetId, req.Data.Content)
	default:
		return errors.ReqTypeErr(req.RequestCommon)
	}

	v := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  req.SeqNum,
		From:    req.To,
		To:      req.From,
		Type:    req.Type,
		Number:  req.Number,
		Message: "success",
		Code:    0,
	}
	out, _ := json.Marshal(v)
	return out
}
