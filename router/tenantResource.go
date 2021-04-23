package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type tenantResource struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewTenantResourceRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &tenantResource{router: router, response: &response.Handler{}}
}

func (r *tenantResource) Init() {
	group := r.router.Group("/tenantResource")
	{
		group.POST("create", r.response.Handler()(api.TenantResource.Create))             // 新建TenantResource
		group.GET("first", r.response.Handler()(api.TenantResource.First))                // 根据ID获取TenantResource
		group.PUT("update", r.response.Handler()(api.TenantResource.Update))              // 更新TenantResource
		group.DELETE("delete", r.response.Handler()(api.TenantResource.Delete))           // 删除TenantResource
		group.DELETE("deletes", r.response.Handler()(api.TenantResource.Deletes))         // 批量删除TenantResource
		group.GET("getList", r.response.Handler()(api.TenantResource.GetList))            // 获取TenantResource列表
	}
}