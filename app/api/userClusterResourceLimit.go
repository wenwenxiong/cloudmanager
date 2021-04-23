package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var UserClusterResourceLimit = new(userClusterResourceLimit)

type userClusterResourceLimit struct{}

// @Tags UserClusterResourceLimit
// @Summary 创建UserClusterResourceLimit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserClusterResourceLimit true "创建UserClusterResourceLimit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userClusterResourceLimit/create [post]
func (a *userClusterResourceLimit) Create(r *ghttp.Request) *response.Response {
	var info model.UserClusterResourceLimit
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.UserClusterResourceLimit.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags UserClusterResourceLimit
// @Summary 用id查询UserClusterResourceLimit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userClusterResourceLimit/first [get]
func (a *userClusterResourceLimit) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.UserClusterResourceLimit.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"userClusterResourceLimit": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags UserClusterResourceLimit
// @Summary 更新UserClusterResourceLimit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserClusterResourceLimit true "userClusterResourceLimit模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userClusterResourceLimit/update [put]
func (a *userClusterResourceLimit) Update(r *ghttp.Request) *response.Response {
	var info model.UserClusterResourceLimit
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.UserClusterResourceLimit.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"userClusterResourceLimit": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags UserClusterResourceLimit
// @Summary 删除UserClusterResourceLimit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "UserClusterResourceLimit模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userClusterResourceLimit/delete [delete]
func (a *userClusterResourceLimit) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.UserClusterResourceLimit.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags UserClusterResourceLimit
// @Summary 批量删除UserClusterResourceLimit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除UserClusterResourceLimit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userClusterResourceLimit/deletes [delete]
func (a *userClusterResourceLimit) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.UserClusterResourceLimit.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags UserClusterResourceLimit
// @Summary 分页获取UserClusterResourceLimit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchUserClusterResourceLimit true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userClusterResourceLimit/getList [get]
func (a *userClusterResourceLimit) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchUserClusterResourceLimit
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.UserClusterResourceLimit.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
