package admin

import (
	"CloudNote/app/system/admin/internal/api"
	"CloudNote/app/system/admin/internal/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()
	// 跨域请求中间件
	s.Use(service.Middleware.CORS)
	apiVersion := g.Cfg().GetString("cloudNote.apiVersion", "/api/v1")
	s.Group(apiVersion, func(group *ghttp.RouterGroup) {

		group.POST("/users", api.User.SignUp) // 用户注册

		group.POST("/user/login", api.User.LogIn) // 用户登录

		// 鉴权中间件
		group.Middleware(service.Middleware.Auth)

		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.PUT("/password", api.User.UpdatePassword) // 修改密码
			group.PUT("/profile", api.User.UpdateProfile)   // 修改用户资料
			group.GET("/profile", api.User.GetProfile)      // 查询用户资料
		})

		group.GET("/notes", api.Undefine)    // 查询笔记列表 (不包含笔记内容)
		group.POST("/notes", api.Undefine)   // 新增笔记
		group.DELETE("/notes", api.Undefine) // 新增笔记

		group.GET("/note", api.Undefine) // 查询笔记详情

		group.Group("/note", func(group *ghttp.RouterGroup) {
			group.PUT("/limit", api.Undefine)    // 设置笔记权限
			group.PUT("/title", api.Undefine)    // 更改笔记标题
			group.PUT("/content", api.Undefine)  // 更改笔记内容
			group.PUT("/category", api.Undefine) // 更改笔记内容
		})

		group.GET("/noteTrash", api.Undefine)    // 查询回收站中的笔记
		group.DELETE("/noteTrash", api.Undefine) // 从回收站中删除
		group.PUT("/noteTrash", api.Undefine)    // 从回收站中恢复

		group.GET("/categories", api.Undefine)    // 查询分类列表
		group.POST("/categories", api.Undefine)   // 新增分类
		group.DELETE("/categories", api.Undefine) // 删除分类

		group.Group("/categories", func(group *ghttp.RouterGroup) {
			group.PUT("/title", api.Undefine)       // 编辑分类名称
			group.PUT("/description", api.Undefine) // 编辑分类描述信息
		})

		group.GET("/plans", api.Plan.QueryList) // 查询待办列表
		group.POST("/plans", api.Plan.Add)      // 新增待办
		group.Group("/plan", func(group *ghttp.RouterGroup) {
			group.PUT("/title", api.Plan.UpdateTitle)     // 编辑待办名称
			group.PUT("/content", api.Plan.UpdateContent) // 编辑待办内容

		})
		group.PUT("/plan", api.Plan.UpdateStatus) // 更新待办状态
		group.DELETE("/plans", api.Plan.Delete)   // 删除待办
	})
}
