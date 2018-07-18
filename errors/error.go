package errors

import (
	"encoding/json"

	"github.com/wst-libs/wstsdk/utils"
)

const (
	JsonParseErr = 4001
)

///
/// Failed to parse json
///
func ParseJsonFailed() []byte {
	comm := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
		From:    "",
		To:      "",
		Type:    "",
		Number:  "",
		Message: "Failed to parse json",
		Code:    JsonParseErr,
	}
	out, _ := json.Marshal(comm)

	return out
}

///
/// Request param is null
///
func RequestParamFiled() []byte {
	comm := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
	}
	out, _ := json.Marshal(comm)
	return out
}
