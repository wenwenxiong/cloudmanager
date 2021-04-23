package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type clusterSpec struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewClusterSpecRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &clusterSpec{router: router, response: &response.Handler{}}
}

func (r *clusterSpec) Init() {
	group := r.router.Group("/clusterSpec")
	{
		group.POST("create", r.response.Handler()(api.ClusterSpec.Create))             // 新建ClusterSpec
		group.GET("first", r.response.Handler()(api.ClusterSpec.First))                // 根据ID获取ClusterSpec
		group.PUT("update", r.response.Handler()(api.ClusterSpec.Update))              // 更新ClusterSpec
		group.DELETE("delete", r.response.Handler()(api.ClusterSpec.Delete))           // 删除ClusterSpec
		group.DELETE("deletes", r.response.Handler()(api.ClusterSpec.Deletes))         // 批量删除ClusterSpec
		group.GET("getList", r.response.Handler()(api.ClusterSpec.GetList))            // 获取ClusterSpec列表
	}
}