package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type credential struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCredentialRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &credential{router: router, response: &response.Handler{}}
}

func (r *credential) Init() {
	group := r.router.Group("/credential")
	{
		group.POST("create", r.response.Handler()(api.Credential.Create))             // 新建Credential
		group.GET("first", r.response.Handler()(api.Credential.First))                // 根据ID获取Credential
		group.PUT("update", r.response.Handler()(api.Credential.Update))              // 更新Credential
		group.DELETE("delete", r.response.Handler()(api.Credential.Delete))           // 删除Credential
		group.DELETE("deletes", r.response.Handler()(api.Credential.Deletes))         // 批量删除Credential
		group.GET("getList", r.response.Handler()(api.Credential.GetList))            // 获取Credential列表
	}
}