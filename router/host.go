package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type host struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewHostRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &host{router: router, response: &response.Handler{}}
}

func (r *host) Init() {
	group := r.router.Group("/host")
	{
		group.POST("create", r.response.Handler()(api.Host.Create))             // 新建Host
		group.GET("first", r.response.Handler()(api.Host.First))                // 根据ID获取Host
		group.PUT("update", r.response.Handler()(api.Host.Update))              // 更新Host
		group.DELETE("delete", r.response.Handler()(api.Host.Delete))           // 删除Host
		group.DELETE("deletes", r.response.Handler()(api.Host.Deletes))         // 批量删除Host
		group.GET("getList", r.response.Handler()(api.Host.GetList))            // 获取Host列表
	}
}