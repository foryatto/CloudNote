package define

import (
	"CloudNote/app/shared"
)

type CategoryAddReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CategoryDeleteReq struct {
	CategoryIds []string `json:"categoryIds"`
}

type CategoryUpdateTitleReq struct {
	Title      string `json:"title"`
	CategoryId string `json:"categoryId"`
}

type CategoryUpdateDescriptionReq struct {
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
}

type CategoryQueryReq struct {
	shared.StandReqParam
	CategoryId  string `json:"categoryId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CategoryQueryResp struct {
	CategoryId  string `json:"categoryId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
