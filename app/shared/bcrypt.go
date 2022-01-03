package shared

import (
	"github.com/gogf/gf/frame/g"
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		g.Log().Line().Warning(err)
		return ""
	}
	return string(hash)
}

// ValidatePassword 验证密码
func ValidatePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
