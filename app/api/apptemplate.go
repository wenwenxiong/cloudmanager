package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Apptemplate = new(apptemplate)

type apptemplate struct{}

// @Tags Apptemplate
// @Summary 创建Apptemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apptemplate true "创建Apptemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /apptemplate/create [post]
func (a *apptemplate) Create(r *ghttp.Request) *response.Response {
	var info model.Apptemplate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.Apptemplate.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags Apptemplate
// @Summary 用id查询Apptemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apptemplate/first [get]
func (a *apptemplate) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.Apptemplate.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"apptemplate": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags Apptemplate
// @Summary 更新Apptemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apptemplate true "apptemplate模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apptemplate/update [put]
func (a *apptemplate) Update(r *ghttp.Request) *response.Response {
	var info model.Apptemplate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.Apptemplate.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"apptemplate": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags Apptemplate
// @Summary 删除Apptemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "Apptemplate模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apptemplate/delete [delete]
func (a *apptemplate) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Apptemplate.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags Apptemplate
// @Summary 批量删除Apptemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除Apptemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apptemplate/deletes [delete]
func (a *apptemplate) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.Apptemplate.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags Apptemplate
// @Summary 分页获取Apptemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApptemplate true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apptemplate/getList [get]
func (a *apptemplate) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchApptemplate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.Apptemplate.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
