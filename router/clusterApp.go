package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type clusterApp struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewClusterAppRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &clusterApp{router: router, response: &response.Handler{}}
}

func (r *clusterApp) Init() {
	group := r.router.Group("/clusterApp")
	{
		group.POST("create", r.response.Handler()(api.ClusterApp.Create))             // 新建ClusterApp
		group.GET("first", r.response.Handler()(api.ClusterApp.First))                // 根据ID获取ClusterApp
		group.PUT("update", r.response.Handler()(api.ClusterApp.Update))              // 更新ClusterApp
		group.DELETE("delete", r.response.Handler()(api.ClusterApp.Delete))           // 删除ClusterApp
		group.DELETE("deletes", r.response.Handler()(api.ClusterApp.Deletes))         // 批量删除ClusterApp
		group.GET("getList", r.response.Handler()(api.ClusterApp.GetList))            // 获取ClusterApp列表
	}
}