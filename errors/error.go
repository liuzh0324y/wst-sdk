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
		Version: "V1.0",
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
