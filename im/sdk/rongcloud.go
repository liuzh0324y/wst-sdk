package sdk

import (
	"fmt"

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

func (s *RCServer) GetTokenFromUser(uid, name, portraitUri string) string {
	ret, _ := rongcloud.User.GetToken(uid, name, portraitUri)
	if ret.Code != 200 {

	}

	fmt.Printf("%v\n", ret)
	return ret.Token
}
