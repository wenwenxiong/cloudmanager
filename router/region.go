package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type region struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewRegionRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &region{router: router, response: &response.Handler{}}
}

func (r *region) Init() {
	group := r.router.Group("/region")
	{
		group.POST("create", r.response.Handler()(api.Region.Create))             // 新建Region
		group.GET("first", r.response.Handler()(api.Region.First))                // 根据ID获取Region
		group.PUT("update", r.response.Handler()(api.Region.Update))              // 更新Region
		group.DELETE("delete", r.response.Handler()(api.Region.Delete))           // 删除Region
		group.DELETE("deletes", r.response.Handler()(api.Region.Deletes))         // 批量删除Region
		group.GET("getList", r.response.Handler()(api.Region.GetList))            // 获取Region列表
	}
}