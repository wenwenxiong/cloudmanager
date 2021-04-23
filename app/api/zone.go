package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Zone = new(zone)

type zone struct{}

// @Tags Zone
// @Summary 创建Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Zone true "创建Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /zone/create [post]
func (a *zone) Create(r *ghttp.Request) *response.Response {
	var info model.Zone
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.Zone.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags Zone
// @Summary 用id查询Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zone/first [get]
func (a *zone) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.Zone.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"zone": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags Zone
// @Summary 更新Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Zone true "zone模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zone/update [put]
func (a *zone) Update(r *ghttp.Request) *response.Response {
	var info model.Zone
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.Zone.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"zone": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags Zone
// @Summary 删除Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "Zone模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zone/delete [delete]
func (a *zone) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Zone.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags Zone
// @Summary 批量删除Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /zone/deletes [delete]
func (a *zone) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.Zone.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags Zone
// @Summary 分页获取Zone列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchZone true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zone/getList [get]
func (a *zone) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchZone
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.Zone.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
