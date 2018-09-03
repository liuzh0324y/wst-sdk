package omigad

import (
	"encoding/json"

	"github.com/wst-libs/wst-sdk/utils"
)

func CommonRequest(body []byte) (*utils.RequestCommon, error) {
	var com utils.RequestCommon
	err := json.Unmarshal(body, &com)
	if err != nil {
		return nil, err
	}
	return &com, nil
}

func PutFileRequest(body []byte) (*ReqPutFile, error) {
	var s ReqPutFile
	err := json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func UpdateFileRequest(body []byte) (*ReqUpdateFile, error) {
	var s ReqUpdateFile
	err := json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
