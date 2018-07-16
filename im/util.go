package im

import "github.com/wst-libs/wst-sdk/utils"

///
/// Create session
///
type CreateSession struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type RequestCreateSession struct {
	utils.RequestCommon
	CreateSession `json:"data"`
}

type ResponseCreateSession struct {
	utils.ResponseCommon
	utils.ID `json:"data"`
}

///
/// Delete session
///
type DelSession struct {
	List []utils.ID `json:"list"`
	Size int64      `json:"size"`
}

type RequestDelSession struct {
	utils.RequestCommon
	DelSession `json:"data"`
}

type ResponseDelSession struct {
	utils.ResponseCommon
}

///
/// Query session information
///
type QuerySession struct {
	List []utils.RoomInfo `json:"list"`
	Size int64            `json:"size"`
}

type ResponseQuerySession struct {
	utils.ResponseCommon
	QuerySession `json:"data"`
}

///
/// Query the list of users of the session
///
type QuerySessionUsers struct {
	List []utils.ID `json:"list"`
	Size int64      `json:"size"`
}

type ResponseQuerySessionUsers struct {
	utils.ResponseCommon
	QuerySessionUsers `json:"data"`
}

///
/// Join the session
///
type JoinSession struct {
}

type RequestJoinSession struct {
	utils.RequestCommon
}

type QuestJoinSession struct {
	utils.ResponseCommon
}

///
/// Quit the session
///
type RequestQuitSession struct {
	utils.RequestCommon
}

type ResponseQuitSession struct {
	utils.ResponseCommon
}

///
/// Registered users
///
type RegisteredUsers struct {
	Name     string `json:"name"`
	Portrait string `json:"portrait"`
	Type     string `json:"type"`
}

type RequestRegisteredUsers struct {
	utils.RequestCommon
	RegisteredUsers `json:"data"`
}

type ResponseRegisteredUsers struct {
	utils.ResponseCommon
	utils.TOKEN `json:"data"`
}

///
/// Send message
///
type SendMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type RequestSendMessage struct {
	utils.RequestCommon
	SendMessage `json:"data"`
}

type ResponseSendMessage struct {
	utils.ResponseCommon
}
