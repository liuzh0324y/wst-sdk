package manager

import "github.com/wst-libs/wst-sdk/utils"

type ResCode struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
	Id   string `json:"id"`
}

type PutFile struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Bucket   string `json:"bucket"`
	Object   string `json:"object"`
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
	Desc     string `json:"description"`
	Version  string `json:"version"`
	Mode     string `json:"model"`
	Content  string `json:"Content"`
	Size     int64  `json:"size"`
	Level    int    `json:"level"`
}

type ReqPutFile struct {
	utils.RequestCommon
	PutFile `json:"data"`
}

type ResPutFile struct {
	utils.ResponseCommon
	utils.ID `json:"data"`
}
