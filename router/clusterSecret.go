package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type clusterSecret struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewClusterSecretRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &clusterSecret{router: router, response: &response.Handler{}}
}

func (r *clusterSecret) Init() {
	group := r.router.Group("/clusterSecret")
	{
		group.POST("create", r.response.Handler()(api.ClusterSecret.Create))             // 新建ClusterSecret
		group.GET("first", r.response.Handler()(api.ClusterSecret.First))                // 根据ID获取ClusterSecret
		group.PUT("update", r.response.Handler()(api.ClusterSecret.Update))              // 更新ClusterSecret
		group.DELETE("delete", r.response.Handler()(api.ClusterSecret.Delete))           // 删除ClusterSecret
		group.DELETE("deletes", r.response.Handler()(api.ClusterSecret.Deletes))         // 批量删除ClusterSecret
		group.GET("getList", r.response.Handler()(api.ClusterSecret.GetList))            // 获取ClusterSecret列表
	}
}