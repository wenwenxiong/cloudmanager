package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type userClusterResourceLimit struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewUserClusterResourceLimitRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &userClusterResourceLimit{router: router, response: &response.Handler{}}
}

func (r *userClusterResourceLimit) Init() {
	group := r.router.Group("/userClusterResourceLimit")
	{
		group.POST("create", r.response.Handler()(api.UserClusterResourceLimit.Create))             // 新建UserClusterResourceLimit
		group.GET("first", r.response.Handler()(api.UserClusterResourceLimit.First))                // 根据ID获取UserClusterResourceLimit
		group.PUT("update", r.response.Handler()(api.UserClusterResourceLimit.Update))              // 更新UserClusterResourceLimit
		group.DELETE("delete", r.response.Handler()(api.UserClusterResourceLimit.Delete))           // 删除UserClusterResourceLimit
		group.DELETE("deletes", r.response.Handler()(api.UserClusterResourceLimit.Deletes))         // 批量删除UserClusterResourceLimit
		group.GET("getList", r.response.Handler()(api.UserClusterResourceLimit.GetList))            // 获取UserClusterResourceLimit列表
	}
}