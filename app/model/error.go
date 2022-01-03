package model

const (
	/*
		cloudNote project error
	*/
	ERR_SIGNUP     = -1000
	ERR_USER_EXIST = -1001

	ERR_LOGIN          = -1010
	ERR_LOGIN_PASSWORD = -1011
	ERR_LOGIN_NOEXIST  = -1012

	ERR_USER_CHANGEPWD = -1020
	/*
		database error code ( -2000 ~ -2009 )
	*/
	ERR_DATABASE_CREATE = -2000
	ERR_DATABASE_READ   = -2001
	ERR_DATABASE_UPDATE = -2002
	ERR_DATABASE_DELETE = -2003
	ERR_DATABASE_COUNT  = -2004
	/*
		http error code ( -2010 ~  )
	*/
	ERR_INVALID_PARAM = -2010
)

var ERR_MAP = map[int]string{
	ERR_SIGNUP:     "sign up failed",
	ERR_USER_EXIST: "user existed",

	ERR_LOGIN:          "login failed",
	ERR_LOGIN_PASSWORD: "login password error",
	ERR_LOGIN_NOEXIST:  "user no exist",

	ERR_USER_CHANGEPWD: "change password failed",

	ERR_DATABASE_CREATE: "database create error",
	ERR_DATABASE_READ:   "database read error",
	ERR_DATABASE_UPDATE: "database update error",
	ERR_DATABASE_DELETE: "database delete error",
	ERR_DATABASE_COUNT:  "database count error",
	ERR_INVALID_PARAM:   "invalid param",
}
