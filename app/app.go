package app

import (
	"CloudNote/app/shared"
	"CloudNote/app/system/admin"
	"github.com/gogf/gf/frame/g"
	_ "github.com/lib/pq"
	"time"
)

func init() {
	// 初始化 jwt Todo
	//shared.JwtUtil.Set(int64(24*time.Hour), "cloudNote", guuid.New().String())
	shared.JwtUtil.Set(int64(24*time.Hour), "cloudNote", "123")
	//初始化 admin
	admin.Init()
}

func Run() {
	s := g.Server()
	s.SetServerRoot("public")
	s.AddStaticPath("/m", "public/client/")
	s.AddStaticPath("/admin", "public/admin/")
	s.Run()
}
