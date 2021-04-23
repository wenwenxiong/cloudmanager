package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type tenant struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewTenantRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &tenant{router: router, response: &response.Handler{}}
}

func (r *tenant) Init() {
	group := r.router.Group("/tenant")
	{
		group.POST("create", r.response.Handler()(api.Tenant.Create))             // 新建Tenant
		group.GET("first", r.response.Handler()(api.Tenant.First))                // 根据ID获取Tenant
		group.PUT("update", r.response.Handler()(api.Tenant.Update))              // 更新Tenant
		group.DELETE("delete", r.response.Handler()(api.Tenant.Delete))           // 删除Tenant
		group.DELETE("deletes", r.response.Handler()(api.Tenant.Deletes))         // 批量删除Tenant
		group.GET("getList", r.response.Handler()(api.Tenant.GetList))            // 获取Tenant列表
	}
}