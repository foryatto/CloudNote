package api

import (
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"CloudNote/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Note = noteApi{}

type noteApi struct{}

func (n *noteApi) Add(r *ghttp.Request) {
	var param *define.NoteAddReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.Add(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_CREATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) UpdateLimit(r *ghttp.Request) {
	var param *define.NoteUpdateLimitReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.UpdateLimit(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) UpdateTitle(r *ghttp.Request) {
	var param *define.NoteUpdateTitleReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.UpdateTitle(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) UpdateContent(r *ghttp.Request) {
	var param *define.NoteUpdateContentReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.UpdateContent(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) UpdateCategory(r *ghttp.Request) {
	var param *define.NoteUpdateCategoryReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.UpdateCategory(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) Delete(r *ghttp.Request) {
	var param *define.NoteDeleteReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.Delete(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_DELETE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) RecoverFromTrash(r *ghttp.Request) {
	var param *define.NoteRecoverReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.RecoverFromTrash(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) DeleteFromTrash(r *ghttp.Request) {
	var param *define.NoteDeleteFromTrashReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	err := service.Note.DeleteFromTrash(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_DELETE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (n *noteApi) BaseQuery(r *ghttp.Request) {
	var param *define.NoteBaseQueryReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	result, err := service.Note.BaseQuery(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, result)
}

func (n *noteApi) DetailQuery(r *ghttp.Request) {
	var param *define.NoteDetailQueryReq
	shared.Parse(r, &param)
	ownerId := r.GetCtxVar("uid").String()
	result, err := service.Note.DetailQuery(ownerId, param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, result)
}

func (n *noteApi) BaseQueryByUserId(r *ghttp.Request) {
	var param *define.NotePublicBaseQueryReq
	shared.Parse(r, &param)
	result, err := service.Note.BaseQueryByUserId(param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, result)
}

func (n *noteApi) DetailQueryByUserId(r *ghttp.Request) {
	var param *define.NotePublicDetailQueryReq
	shared.Parse(r, &param)
	result, err := service.Note.DetailQueryByUserId(param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, result)
}
