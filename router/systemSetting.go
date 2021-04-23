package router

import (
	"cloudmanager/app/api"
	"cloudmanager/library/response"
	"cloudmanager/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type systemSetting struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewSystemSettingRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &systemSetting{router: router, response: &response.Handler{}}
}

func (r *systemSetting) Init() {
	group := r.router.Group("/systemSetting")
	{
		group.POST("create", r.response.Handler()(api.SystemSetting.Create))             // 新建SystemSetting
		group.GET("first", r.response.Handler()(api.SystemSetting.First))                // 根据ID获取SystemSetting
		group.PUT("update", r.response.Handler()(api.SystemSetting.Update))              // 更新SystemSetting
		group.DELETE("delete", r.response.Handler()(api.SystemSetting.Delete))           // 删除SystemSetting
		group.DELETE("deletes", r.response.Handler()(api.SystemSetting.Deletes))         // 批量删除SystemSetting
		group.GET("getList", r.response.Handler()(api.SystemSetting.GetList))            // 获取SystemSetting列表
	}
}