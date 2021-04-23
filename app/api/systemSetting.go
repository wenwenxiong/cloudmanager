package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var SystemSetting = new(systemSetting)

type systemSetting struct{}

// @Tags SystemSetting
// @Summary 创建SystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemSetting true "创建SystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /systemSetting/create [post]
func (a *systemSetting) Create(r *ghttp.Request) *response.Response {
	var info model.SystemSetting
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.SystemSetting.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags SystemSetting
// @Summary 用id查询SystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /systemSetting/first [get]
func (a *systemSetting) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.SystemSetting.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"systemSetting": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags SystemSetting
// @Summary 更新SystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemSetting true "systemSetting模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /systemSetting/update [put]
func (a *systemSetting) Update(r *ghttp.Request) *response.Response {
	var info model.SystemSetting
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.SystemSetting.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"systemSetting": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags SystemSetting
// @Summary 删除SystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "SystemSetting模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemSetting/delete [delete]
func (a *systemSetting) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.SystemSetting.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags SystemSetting
// @Summary 批量删除SystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除SystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /systemSetting/deletes [delete]
func (a *systemSetting) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.SystemSetting.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags SystemSetting
// @Summary 分页获取SystemSetting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchSystemSetting true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemSetting/getList [get]
func (a *systemSetting) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchSystemSetting
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.SystemSetting.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
