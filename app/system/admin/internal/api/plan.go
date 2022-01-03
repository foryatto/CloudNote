package api

import (
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"CloudNote/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Plan = planApi{}

type planApi struct{}

func (p *planApi) QueryList(r *ghttp.Request) {
	var param *define.PlanQueryReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	result, err := service.Plan.QueryList(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, result)
}

func (p *planApi) Add(r *ghttp.Request) {
	var param *define.PlanAddReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Plan.Add(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_CREATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (p *planApi) UpdateTitle(r *ghttp.Request) {
	var param *define.PlanUpdateTitleReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Plan.UpdateTitle(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (p *planApi) UpdateContent(r *ghttp.Request) {
	var param *define.PlanUpdateContentReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Plan.UpdateContent(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (p *planApi) UpdateStatus(r *ghttp.Request) {
	var param *define.PlanUpdateStatusReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Plan.UpdateStatus(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (p *planApi) Delete(r *ghttp.Request) {
	var param *define.PlanDeleteReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Plan.Delete(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_DELETE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}
