package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type tenantMember struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewTenantMemberRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &tenantMember{router: router, response: &response.Handler{}}
}

func (r *tenantMember) Init() {
	group := r.router.Group("/tenantMember")
	{
		group.POST("create", r.response.Handler()(api.TenantMember.Create))             // 新建TenantMember
		group.GET("first", r.response.Handler()(api.TenantMember.First))                // 根据ID获取TenantMember
		group.PUT("update", r.response.Handler()(api.TenantMember.Update))              // 更新TenantMember
		group.DELETE("delete", r.response.Handler()(api.TenantMember.Delete))           // 删除TenantMember
		group.DELETE("deletes", r.response.Handler()(api.TenantMember.Deletes))         // 批量删除TenantMember
		group.GET("getList", r.response.Handler()(api.TenantMember.GetList))            // 获取TenantMember列表
	}
}