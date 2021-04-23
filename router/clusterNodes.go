package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type clusterNodes struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewClusterNodesRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &clusterNodes{router: router, response: &response.Handler{}}
}

func (r *clusterNodes) Init() {
	group := r.router.Group("/clusterNodes")
	{
		group.POST("create", r.response.Handler()(api.ClusterNodes.Create))             // 新建ClusterNodes
		group.GET("first", r.response.Handler()(api.ClusterNodes.First))                // 根据ID获取ClusterNodes
		group.PUT("update", r.response.Handler()(api.ClusterNodes.Update))              // 更新ClusterNodes
		group.DELETE("delete", r.response.Handler()(api.ClusterNodes.Delete))           // 删除ClusterNodes
		group.DELETE("deletes", r.response.Handler()(api.ClusterNodes.Deletes))         // 批量删除ClusterNodes
		group.GET("getList", r.response.Handler()(api.ClusterNodes.GetList))            // 获取ClusterNodes列表
	}
}