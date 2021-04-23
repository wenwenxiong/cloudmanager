package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type user struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewUserRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &user{router: router, response: &response.Handler{}}
}

func (r *user) Init() {
	group := r.router.Group("/user")
	{
		group.POST("create", r.response.Handler()(api.User.Create))             // 新建User
		group.GET("first", r.response.Handler()(api.User.First))                // 根据ID获取User
		group.PUT("update", r.response.Handler()(api.User.Update))              // 更新User
		group.DELETE("delete", r.response.Handler()(api.User.Delete))           // 删除User
		group.DELETE("deletes", r.response.Handler()(api.User.Deletes))         // 批量删除User
		group.GET("getList", r.response.Handler()(api.User.GetList))            // 获取User列表
	}
}