package define

import (
	"CloudNote/app/shared"
	"github.com/gogf/gf/os/gtime"
)

type PlanAddReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PlanQueryReq struct {
	shared.StandReqParam
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
	PlanId    string `json:"planId"`
}

type PlanQueryResp struct {
	PlanId    string      `json:"planId"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
	Completed bool        `json:"completed"`
}

type PlanUpdateTitleReq struct {
	Title  string `json:"title"`
	PlanId string `json:"planId"`
}

type PlanUpdateContentReq struct {
	Content string `json:"content"`
	PlanId  string `json:"planId"`
}

type PlanUpdateStatusReq struct {
	Completed string `json:"completed"`
	PlanId    string `json:"planId"`
}

type PlanDeleteReq struct {
	PlanIds []string `json:"ids"`
}
