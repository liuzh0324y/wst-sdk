package sdk

import (
	"errors"
	"log"

	"github.com/rongcloud/server-sdk-go/RCServerSDK"
)

var rongcloud *rcserversdk.RongCloud

var app = "z3v5yqkbz1xg0"
var sec = "NVx79htIGLm"

///
/// Room info
///
type RoomInfo struct {
	Id   string
	Name string
	Time string
}

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

func (s *RCServer) GetChatRoomById(id string) (RoomInfo, error) {
	ret, err := rongcloud.Chatroom.Query([]string{id})
	if err != nil {
		log.Println("GetChatRoomById Error: ", err.Error())
		return RoomInfo{}, err
	}
	if len(ret.ChatRooms) == 0 {
		return RoomInfo{}, errors.New("Not found session info")
	}
	return RoomInfo{
		Id:   ret.ChatRooms[0].ChrmId,
		Name: ret.ChatRooms[0].Name,
		Time: ret.ChatRooms[0].Time,
	}, nil
}

func (s *RCServer) GetUsersByRoomId(id string) ([]string, error) {
	ret, err := rongcloud.Chatroom.QueryUser(id, "500", "1")
	if err != nil {
		log.Println("GetUsersByRoomId Error: ", err.Error())
		return []string{}, err
	}
	if len(ret.Users) == 0 {
		log.Println("GetUsersByRoomId Error: Not found users in room")
		return []string{}, errors.New("Not found users in room")
	}
	var users []string
	for i, user := range ret.Users {
		users[i] = user.Id
	}
	return users, nil
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
