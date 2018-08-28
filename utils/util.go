package utils

const (
	Version = "V1.0"
)

///
/// Common request structure
///
type RequestCommon struct {
	Version string `json:"version"`
	SeqNum  int64  `json:"seqnum"`
	From    string `json:"from"`
	To      string `json:"to"`
	Type    string `json:"type"`
	Number  string `json:"number"`
	Uid     string `json:"uid"`
}

///
/// Common response structure
///
type ResponseCommon struct {
	Version string `json:"version"`
	SeqNum  int64  `json:"seqnum"`
	From    string `json:"from"`
	To      string `json:"to"`
	Type    string `json:"type"`
	Number  string `json:"number"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type ID struct {
	Id string `json:"id"`
}

type Url struct {
	Url string `json:"url"`
}

type TOKEN struct {
	Token string `json:"token"`
}

type RoomInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CreateTime  string `json:"createtime"`
	UpdateTime  string `json:"updatetime"`
}

func GetSeqNum() int64 {
	return 0
}
