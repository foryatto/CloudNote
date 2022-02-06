package define

import (
	"CloudNote/app/shared"
	"github.com/gogf/gf/os/gtime"
)

type Note struct {
	NoteId      string      `orm:"note_id"      json:"noteId"`
	Title       string      `orm:"title"        json:"title"`
	Content     string      `orm:"content"      json:"content"`
	CreatedAt   *gtime.Time `orm:"created_at"   json:"createdAt"`
	UpdatedAt   *gtime.Time `orm:"updated_at"   json:"updatedAt"`
	OwnerId     string      `orm:"owner_id"     json:"ownerId"`
	Shared      bool        `orm:"shared"       json:"shared"`
	Trash       bool        `orm:"trash"        json:"trash"`
	DeletedTime *gtime.Time `orm:"deleted_time" json:"deletedTime"`

	Category Category `json:"category"`
}

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
	Title     string      `json:"title"`
	NoteId    string      `json:"noteId"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at"`
	Category  Category    `json:"category"`
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
