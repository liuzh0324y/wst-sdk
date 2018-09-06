package omigad

import "github.com/wst-libs/wst-sdk/utils"

const (
	defaultBucket = "llvision"
)

// PutFile request structure
type ReqPutFile struct {
	Data struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Path   string `json:"path"`
		Url    string `json:"url"`
		Size   int64  `json:"size"`
		Bucket string `json:"bucket"`
		Object string `json:"object"`
		Desc   string `json:"description"`
	} `json:"data"`
}

// PutFile response structure
type ResPutFile struct {
	utils.ResponseCommon
	utils.ID `json:"data"`
}

// GetFile response structure
type GetFileData struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Url    string `json:"url"`
	Size   string `json:"size"`
	Bucket string `json:"bucket"`
	Object string `json:"object"`
	Desc   string `json:"description"`
}
type ResGetFile struct {
	utils.ResponseCommon
	GetFileData `json:"data"`
}

// UpdateFile request structure
type ReqUpdateFile struct {
	Data struct {
		Object string `json:"object"`
		Desc   string `json:"description"`
	}
}

// UpdateFile response structure
type ResUpdateFile struct {
	utils.ResponseCommon
}

type ResDeleteFile struct {
	utils.ResponseCommon
}

type ReqDeleteFileList struct {
}

type ResDeleteFileList struct {
	utils.ResponseCommon
}

type PostInfo struct {
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
}

type ResUpdateFileToP struct {
	utils.RequestCommon
	PostInfo `json:"data"`
}

type PutFileInfo struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type ResGetUrlForFile struct {
	utils.ResponseCommon
	PutFileInfo `json:"data"`
}

type FileInfo struct {
	Id       string
	FilePath string
	FileName string
	FileType string
	Channel  string
	SeqNum   int64
	TmpPath  string
	Bucket   string
	Object   string
}
