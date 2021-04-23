package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type cluster struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewClusterRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &cluster{router: router, response: &response.Handler{}}
}

func (r *cluster) Init() {
	group := r.router.Group("/cluster")
	{
		group.POST("create", r.response.Handler()(api.Cluster.Create))             // 新建Cluster
		group.GET("first", r.response.Handler()(api.Cluster.First))                // 根据ID获取Cluster
		group.PUT("update", r.response.Handler()(api.Cluster.Update))              // 更新Cluster
		group.DELETE("delete", r.response.Handler()(api.Cluster.Delete))           // 删除Cluster
		group.DELETE("deletes", r.response.Handler()(api.Cluster.Deletes))         // 批量删除Cluster
		group.GET("getList", r.response.Handler()(api.Cluster.GetList))            // 获取Cluster列表
	}
}