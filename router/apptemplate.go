package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type apptemplate struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewApptemplateRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &apptemplate{router: router, response: &response.Handler{}}
}

func (r *apptemplate) Init() {
	group := r.router.Group("/apptemplate")
	{
		group.POST("create", r.response.Handler()(api.Apptemplate.Create))             // 新建Apptemplate
		group.GET("first", r.response.Handler()(api.Apptemplate.First))                // 根据ID获取Apptemplate
		group.PUT("update", r.response.Handler()(api.Apptemplate.Update))              // 更新Apptemplate
		group.DELETE("delete", r.response.Handler()(api.Apptemplate.Delete))           // 删除Apptemplate
		group.DELETE("deletes", r.response.Handler()(api.Apptemplate.Deletes))         // 批量删除Apptemplate
		group.GET("getList", r.response.Handler()(api.Apptemplate.GetList))            // 获取Apptemplate列表
	}
}