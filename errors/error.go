package errors

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/wst-libs/wstsdk/utils"
)

const (
	JsonParseErr     = 4001 // error json format
	VersionErr       = 4002 // unknown version
	SeqnumErr        = 4003 // seqnum is null
	FromErr          = 4004 // unknown from address
	ToErr            = 4005 // unknown to address
	TypeErr          = 4006 // unknown service type
	UidErr           = 4007 // err uid format or null
	UserTokenErr     = 4008 // failed to get the user token
	CreateSessionErr = 4009 // failed to create the session by session id
	DeleteSessionErr = 4010 // failed to delete the session by session id
)

type CodeResult struct {
	utils.ResponseCommon
}

// Failed to parse json
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

// Parse common request param
func IsCommonErr(req utils.RequestCommon) *CodeResult {
	if 0 != strings.Compare(utils.Version, req.Version) {
		return &CodeResult{
			utils.ResponseCommon{
				Code:    VersionErr,
				Message: "unknown version",
				Version: req.Version,
				SeqNum:  req.SeqNum,
				From:    req.From,
				To:      req.To,
				Type:    req.Type,
				Number:  req.Number,
			},
		}
	}

	if len(req.Uid) <= 0 {
		return &CodeResult{
			utils.ResponseCommon{
				Code:    UidErr,
				Message: "source uid too short",
				Version: req.Version,
				SeqNum:  req.SeqNum,
				From:    req.From,
				To:      req.To,
				Type:    req.Type,
				Number:  req.Number,
			},
		}

	}

	if len(req.Uid) > 32 {
		return &CodeResult{
			utils.ResponseCommon{
				Code:    UidErr,
				Message: "source uid too long",
				Version: req.Version,
				SeqNum:  req.SeqNum,
				From:    req.From,
				To:      req.To,
				Type:    req.Type,
				Number:  req.Number,
			},
		}

	}

	return &CodeResult{
		utils.ResponseCommon{
			Code:    0,
			Message: "successful",
			Version: req.Version,
			SeqNum:  req.SeqNum,
			From:    req.From,
			To:      req.To,
			Type:    req.Type,
			Number:  req.Number,
		},
	}
}

// Verify that the uid conforms to the rule
func VerifyUid(uid string) error {
	if len(uid) <= 0 {
		return errors.New("uid too short")
	}

	if len(uid) > 32 {
		return errors.New("uid too long")
	}
	return nil
}

// Failed to Rongcloud
func ImplementErr(code int64, req utils.RequestCommon, msg string) []byte {
	out := utils.ResponseCommon{
		Code:    code,
		Message: msg,
		Version: req.Version,
		SeqNum:  req.SeqNum,
		From:    req.From,
		To:      req.To,
		Type:    req.Type,
		Number:  req.Number,
	}
	ret, _ := json.Marshal(out)
	return ret
}

// Request param is null
func RequestParamFiled() []byte {
	comm := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
	}
	out, _ := json.Marshal(comm)
	return out
}
