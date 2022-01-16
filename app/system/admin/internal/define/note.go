package define

import "CloudNote/app/shared"

type NoteAddReq struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryId string `json:"categoryId"`
	Shared     bool   `json:"shared"`
}

type NoteUpdateLimitReq struct {
	Shared bool   `json:"shared"`
	NoteId string `json:"noteId"`
}

type NoteUpdateTitleReq struct {
	Title  string `json:"title"`
	NoteId string `json:"noteId"`
}

type NoteUpdateContentReq struct {
	Content string `json:"content"`
	NoteId  string `json:"noteId"`
}

type NoteUpdateCategoryReq struct {
	CategoryId string `json:"categoryId"`
	NoteId     string `json:"noteId"`
}

type NoteDeleteReq struct {
	NoteIds []string `json:"ids"`
}

type NoteRecoverReq struct {
	NoteIds []string `json:"ids"`
}

type NoteDeleteFromTrashReq struct {
	NoteIds []string `json:"ids"`
}

type NoteBaseQueryReq struct {
	shared.StandReqParam
	Title   string `json:"title"`
	Content string `json:"content"`
	Shared  bool   `json:"shared"`
	NoteId  string `json:"noteId"`
	Trash   bool   `json:"trash"`
}

type NoteBaseQueryResp struct {
	Title  string `json:"title"`
	NoteId string `json:"noteId"`
}

type NoteDetailQueryReq struct {
	NoteId string `json:"noteId"`
}

type NotePublicBaseQueryReq struct {
	UserId string `json:"userId"`
}

type NotePublicDetailQueryReq struct {
	NoteId string `json:"noteId"`
	UserId string `json:"userId"`
}
