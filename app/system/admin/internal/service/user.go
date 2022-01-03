package service

import (
	"CloudNote/app/dao"
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/guuid"
)

var User = userService{}

type userService struct{}

// SignUp 用户注册
func (u *userService) SignUp(param *define.UserSignUpReq) error {
	var result *model.User
	err := dao.User.Ctx(context.TODO()).Fields(dao.User.Columns.Email).
		Scan(&result, dao.User.Columns.Email, param.Email)
	if err != nil {
		g.Log().Line().Warning(err)
		return shared.NewError(model.ERR_SIGNUP)
	}
	if result != nil {
		return shared.NewError(model.ERR_USER_EXIST)
	}
	cryptPwd := shared.HashAndSalt([]byte(param.Password))
	if cryptPwd == "" {
		return shared.NewError(model.ERR_SIGNUP)
	}
	param.Password = cryptPwd
	var saveData *model.User
	err = gconv.Scan(param, &saveData)
	if err != nil {
		g.Log().Line().Warning(err)
		return shared.NewError(model.ERR_SIGNUP)
	}
	saveData.UserId = guuid.New().String()
	_, err = dao.User.Ctx(context.TODO()).Insert(saveData)
	if err != nil {
		g.Log().Line().Warning(err)
		return shared.NewError(model.ERR_SIGNUP)
	}
	return nil
}

func (u *userService) LogIn(param *define.UserLogInReq) (string, error) {
	var result *model.User
	err := dao.User.Ctx(context.TODO()).Fields(dao.User.Columns.Password, dao.User.Columns.UserId).
		Scan(&result, dao.User.Columns.Email, param.Email)
	if err != nil {
		g.Log().Line().Warning(err)
		return "", shared.NewError(model.ERR_LOGIN)
	}
	if result == nil {
		return "", shared.NewError(model.ERR_LOGIN_NOEXIST)
	}
	confirmResult := shared.ValidatePassword(result.Password, []byte(param.Password))
	if !confirmResult {
		return "", shared.NewError(model.ERR_LOGIN_PASSWORD)
	}

	token, err := shared.JwtUtil.Create(result.UserId)
	if err != nil {
		g.Log().Line().Warning(err)
		return "", shared.NewError(model.ERR_LOGIN)
	}
	return token, nil
}

func (u *userService) GetProfile(uid string) (*define.UserProfileResp, error) {
	var result *define.UserProfileResp
	err := dao.User.Ctx(context.TODO()).FieldsEx(dao.User.Columns.Password, dao.User.Columns.UserId).
		Scan(&result, dao.User.Columns.UserId, uid)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	return result, nil
}

func (u *userService) UpdateProfile(uid string, param *define.UserUpdateProfileReq) error {
	_, err := dao.User.Ctx(context.TODO()).Update(g.Map{
		dao.User.Columns.FullName: param.FullName,
	}, dao.User.Columns.UserId, uid)
	if err != nil {
		g.Log().Warning(err)
		return err
	}
	return nil
}

func (u *userService) ChangePwd(uid string, param *define.UserChangePwdReq) error {
	newPwdHashed := shared.HashAndSalt([]byte(param.NewPwd))
	if newPwdHashed == "" {
		return shared.NewError(model.ERR_USER_CHANGEPWD)
	}

	pwdResult, err := dao.User.Ctx(context.TODO()).Fields(dao.User.Columns.Password).Where(g.Map{
		dao.User.Columns.UserId: uid,
	}).One()
	if err != nil {
		g.Log().Line().Warning(err)
		return shared.NewError(model.ERR_USER_CHANGEPWD)
	}
	if pwdResult == nil {
		return shared.NewError(model.ERR_USER_CHANGEPWD)
	}
	oldPwdInDb := pwdResult[dao.User.Columns.Password].String()
	//g.Log().Line().Debug(oldPwdInDb)

	confirmPwd := shared.ValidatePassword(oldPwdInDb, []byte(param.OldPwd))
	if !confirmPwd {
		return shared.NewError(model.ERR_USER_CHANGEPWD)
	}

	_, err = dao.User.Ctx(context.TODO()).Where(g.Map{
		dao.User.Columns.UserId: uid,
	}).Update(g.Map{dao.User.Columns.Password: newPwdHashed})
	if err != nil {
		g.Log().Line().Warning(err)
		return shared.NewError(model.ERR_DATABASE_UPDATE)
	}
	return nil
}
