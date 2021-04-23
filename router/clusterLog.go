package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type clusterLog struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewClusterLogRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &clusterLog{router: router, response: &response.Handler{}}
}

func (r *clusterLog) Init() {
	group := r.router.Group("/clusterLog")
	{
		group.POST("create", r.response.Handler()(api.ClusterLog.Create))             // 新建ClusterLog
		group.GET("first", r.response.Handler()(api.ClusterLog.First))                // 根据ID获取ClusterLog
		group.PUT("update", r.response.Handler()(api.ClusterLog.Update))              // 更新ClusterLog
		group.DELETE("delete", r.response.Handler()(api.ClusterLog.Delete))           // 删除ClusterLog
		group.DELETE("deletes", r.response.Handler()(api.ClusterLog.Deletes))         // 批量删除ClusterLog
		group.GET("getList", r.response.Handler()(api.ClusterLog.GetList))            // 获取ClusterLog列表
	}
}