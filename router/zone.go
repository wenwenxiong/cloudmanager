package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type zone struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewZoneRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &zone{router: router, response: &response.Handler{}}
}

func (r *zone) Init() {
	group := r.router.Group("/zone")
	{
		group.POST("create", r.response.Handler()(api.Zone.Create))             // 新建Zone
		group.GET("first", r.response.Handler()(api.Zone.First))                // 根据ID获取Zone
		group.PUT("update", r.response.Handler()(api.Zone.Update))              // 更新Zone
		group.DELETE("delete", r.response.Handler()(api.Zone.Delete))           // 删除Zone
		group.DELETE("deletes", r.response.Handler()(api.Zone.Deletes))         // 批量删除Zone
		group.GET("getList", r.response.Handler()(api.Zone.GetList))            // 获取Zone列表
	}
}