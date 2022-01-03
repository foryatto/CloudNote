package api

import "github.com/gogf/gf/net/ghttp"

func Undefine(r *ghttp.Request) {
	r.Response.WriteJson("{code:-1}")
}
