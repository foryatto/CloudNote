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
	Completed int   `json:"completed"`  // 0表示查询所有, 1表示完成, 2表示未完成
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
