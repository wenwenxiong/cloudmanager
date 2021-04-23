package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type clusterStatus struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewClusterStatusRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &clusterStatus{router: router, response: &response.Handler{}}
}

func (r *clusterStatus) Init() {
	group := r.router.Group("/clusterStatus")
	{
		group.POST("create", r.response.Handler()(api.ClusterStatus.Create))             // 新建ClusterStatus
		group.GET("first", r.response.Handler()(api.ClusterStatus.First))                // 根据ID获取ClusterStatus
		group.PUT("update", r.response.Handler()(api.ClusterStatus.Update))              // 更新ClusterStatus
		group.DELETE("delete", r.response.Handler()(api.ClusterStatus.Delete))           // 删除ClusterStatus
		group.DELETE("deletes", r.response.Handler()(api.ClusterStatus.Deletes))         // 批量删除ClusterStatus
		group.GET("getList", r.response.Handler()(api.ClusterStatus.GetList))            // 获取ClusterStatus列表
	}
}