package service

import (
	"CloudNote/app/shared"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"strings"
)

var Middleware = middleware{}

type middleware struct {
}

// CORS 允许跨域请求中间件
func (m *middleware) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	//corsOptions.AllowDomain = []string{"localhost", "127.0.0.1"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

// Auth 鉴权中间件
func (m *middleware) Auth(r *ghttp.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	str := strings.Split(token, " ")
	if len(str) < 2 {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	uid, err := shared.JwtUtil.Parse(str[1])
	if err != nil {
		g.Log().Line().Warning(err)
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	//g.Log().Line().Debug(uid)
	r.SetCtxVar("uid", uid)
	r.Middleware.Next()
}
