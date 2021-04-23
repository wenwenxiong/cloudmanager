package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type volume struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewVolumeRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &volume{router: router, response: &response.Handler{}}
}

func (r *volume) Init() {
	group := r.router.Group("/volume")
	{
		group.POST("create", r.response.Handler()(api.Volume.Create))             // 新建Volume
		group.GET("first", r.response.Handler()(api.Volume.First))                // 根据ID获取Volume
		group.PUT("update", r.response.Handler()(api.Volume.Update))              // 更新Volume
		group.DELETE("delete", r.response.Handler()(api.Volume.Delete))           // 删除Volume
		group.DELETE("deletes", r.response.Handler()(api.Volume.Deletes))         // 批量删除Volume
		group.GET("getList", r.response.Handler()(api.Volume.GetList))            // 获取Volume列表
	}
}