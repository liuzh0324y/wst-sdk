package sdk

import (
	"log"

	"github.com/rongcloud/server-sdk-go/RCServerSDK"
)

var rongcloud *rcserversdk.RongCloud

var app = "z3v5yqkbz1xg0"
var sec = "NVx79htIGLm"

type RCServer struct {
}

func NewRCServer() *RCServer {
	rongcloud = rcserversdk.CreateRongCloud(app, sec)
	return &RCServer{}
}

func (s *RCServer) GetTokenFromUser(uid, name, portraitUri string) (string, error) {
	ret, err := rongcloud.User.GetToken(uid, name, portraitUri)
	if err != nil {
		log.Println("GetToken Error: ", err.Error())
		return "", err
	}

	log.Printf("%v\n", ret)
	return ret.Token, nil
}

func (s *RCServer) CreateChatRoom(id, name string) error {
	log.Println("id:", id)
	var roominfo = rcserversdk.ChatRoomInfo{Id: id, Name: name}
	ret, err := rongcloud.Chatroom.Create([]rcserversdk.ChatRoomInfo{roominfo})

	if err != nil {
		log.Println("CreateChatRoom Error: ", err.Error())
		return err
	}
	log.Printf("%v\n", ret)
	return nil
}

func (s *RCServer) DeleteChatRoom(ids []string) error {
	for v := range ids {
		log.Println(v)
	}
	ret, err := rongcloud.Chatroom.Destroy(ids)
	if err != nil {
		log.Println("DeleteChatRoom Error: ", err.Error())
		return err
	}
	log.Printf("%v\n", ret)
	return nil
}

func (s *RCServer) GetChatRoomById(ids []string) {
	ret, _ := rongcloud.Chatroom.Query(ids)
	if ret.Code != 200 {

	}
}

func (s *RCServer) GetUsersByRoomId(id string) {
	ret, _ := rongcloud.Chatroom.QueryUser(id, "500", "1")
	if ret.Code != 200 {

	}
}

func (s *RCServer) JoinRoomByUserId(uid, sid string) {
	rongcloud.Chatroom.Join([]string{uid}, sid)
}

func (s *RCServer) ExitRoomByUserId(uid, sid string) {

}

func (s *RCServer) SendMsgUserToUsers(formId, toId, content string) {
	var vmsg rcserversdk.VoiceMessage
	vmsg.Content = content

	ret, err := rongcloud.Message.PublishPrivate(formId, []string{toId}, vmsg, content, "", "", 0, 1, 1, 1)
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	if ret.Code != 200 {

	}
	log.Printf("response: %v\n", ret)
}

func (s *RCServer) SendMsgUserToSession(formId, toSession, content string) {

	msg := rcserversdk.TxtMessage{
		Content: content,
	}
	msg.SetType("RC:TxtMsg")

	ret, err := rongcloud.Message.PublishChatroom(formId, []string{toSession}, msg)
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	if ret.Code != 200 {

	}
	log.Printf("response: %v\n", ret)
}
