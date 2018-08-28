package manager

import "github.com/wst-libs/wst-sdk/utils"

// Reture result
type ResCode struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
	Id   string `json:"id"`
}

// create and upload record information
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
	Status   int    `json:"status"`
}

// request information for create or upload record
type ReqPutFile struct {
	utils.RequestCommon
	PutFile `json:"data"`
}

// response information for create or upload record
type ResPutFile struct {
	utils.ResponseCommon
	utils.ID `json:"data"`
}
