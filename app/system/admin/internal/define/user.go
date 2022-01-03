package define

import "github.com/gogf/gf/os/gtime"

type UserBaseInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpReq struct {
	UserBaseInfo
	FullName string `json:"fullName"`
}

type UserLogInReq struct {
	UserBaseInfo
}

type UserProfileResp struct {
	Email     string      `json:"email"`
	FullName  string      `json:"fullName"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

type UserUpdateProfileReq struct {
	FullName string `json:"fullName"`
}

type UserChangePwdReq struct {
	OldPwd string `json:"oldPwd"`
	NewPwd string `json:"newPwd"`
}
