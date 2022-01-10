package api

import (
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"CloudNote/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Category = categoryApi{}

type categoryApi struct{}

func (c *categoryApi) Add(r *ghttp.Request) {
	var param *define.CategoryAddReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Category.Add(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_CREATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (c *categoryApi) QueryList(r *ghttp.Request) {
	var param *define.CategoryQueryReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	result, err := service.Category.QueryList(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, result)
}

func (c *categoryApi) Delete(r *ghttp.Request) {
	var param *define.CategoryDeleteReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Category.Delete(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_DELETE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (c *categoryApi) UpdateTitle(r *ghttp.Request) {
	var param *define.CategoryUpdateTitleReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Category.UpdateTitle(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (c *categoryApi) UpdateDescription(r *ghttp.Request) {
	var param *define.CategoryUpdateDescriptionReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Category.UpdateDescription(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}
