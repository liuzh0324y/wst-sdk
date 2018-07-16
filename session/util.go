package session

import "github.com/wst-libs/wst-sdk/utils"

///
/// Create the session
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
/// Delete the session
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
/// Query the session
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
/// Registered users
///
type RequestRegisteredUsers struct {
	utils.RequestCommon
}

type ResponseRegisteredUsers struct {
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
