package api

import (
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"CloudNote/app/system/admin/internal/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct{}

func (u *userApi) SignUp(r *ghttp.Request) {
	var param *define.UserSignUpReq
	shared.Parse(r, &param)
	err := service.User.SignUp(param)
	if err != nil {
		shared.ResponseError(r, gerror.Code(err).Code(), err.Error())
	}
	shared.Response(r, nil)
}

func (u *userApi) LogIn(r *ghttp.Request) {
	var param *define.UserLogInReq
	shared.Parse(r, &param)
	token, err := service.User.LogIn(param)
	if err != nil {
		shared.ResponseError(r, gerror.Code(err).Code(), err.Error())
	}
	shared.Response(r, token)
}

func (u *userApi) GetProfile(r *ghttp.Request) {
	result, err := service.User.GetProfile(r.GetCtxVar("uid").String())
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, result)
}

func (u *userApi) UpdateProfile(r *ghttp.Request) {
	var param *define.UserUpdateProfileReq
	shared.Parse(r, &param)
	err := service.User.UpdateProfile(r.GetCtxVar("uid").String(), param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}

func (u *userApi) UpdatePassword(r *ghttp.Request) {
	var param *define.UserChangePwdReq
	shared.Parse(r, &param)
	err := service.User.ChangePwd(r.GetCtxVar("uid").String(), param)
	if err != nil {
		errCode := model.ERR_DATABASE_UPDATE
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, nil)
}
