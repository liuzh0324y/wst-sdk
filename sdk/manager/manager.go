//
// SDK for the file manager model
//
package manager

import "github.com/wst-libs/wst-sdk/utils"

func Add(url string) ResCode {
	c := &Client{}
	req := &ReqPutFile{
		utils.RequestCommon{
			Version: utils.Version,
			SeqNum:  1,
			From:    "omigad",
			To:      "client",
			Type:    "omigad",
			Number:  "XXXX-XXXX-XXXX-XXXX",
			Uid:     "unknown",
		},
		PutFile{
			Name:     "",
			Type:     "",
			Url:      "",
			Key:      "",
			Secret:   "",
			Bucket:   "",
			Object:   "",
			Region:   "",
			Endpoint: "",
			Desc:     "",
			Version:  "",
			Mode:     "",
			Content:  "",
			Size:     1,
			Level:    1,
		},
	}
	return c.Add(url, req)
}

func Del(url string) {
	c := &Client{}
	c.Del()
}

func Update(url string, status int) {
	c := &Client{}
	req := &ReqPutFile{
		utils.RequestCommon{
			Version: utils.Version,
			SeqNum:  1,
			From:    "omigad",
			To:      "client",
			Type:    "omigad",
			Number:  "XXXX-XXXX-XXXX-XXXX",
			Uid:     "unknown",
		},
		PutFile{
			Name:     "",
			Type:     "",
			Url:      "",
			Key:      "",
			Secret:   "",
			Bucket:   "",
			Object:   "",
			Region:   "",
			Endpoint: "",
			Desc:     "",
			Version:  "",
			Mode:     "",
			Content:  "",
			Size:     1,
			Level:    1,
			Status:   status,
		},
	}
	c.Update(url, req)
}

func Get(url string) {
	c := &Client{}
	c.Get()
}
