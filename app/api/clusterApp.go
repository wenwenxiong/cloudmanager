package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ClusterApp = new(clusterApp)

type clusterApp struct{}

// @Tags ClusterApp
// @Summary 创建ClusterApp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterApp true "创建ClusterApp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clusterApp/create [post]
func (a *clusterApp) Create(r *ghttp.Request) *response.Response {
	var info model.ClusterApp
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.ClusterApp.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags ClusterApp
// @Summary 用id查询ClusterApp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clusterApp/first [get]
func (a *clusterApp) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.ClusterApp.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"clusterApp": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags ClusterApp
// @Summary 更新ClusterApp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterApp true "clusterApp模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clusterApp/update [put]
func (a *clusterApp) Update(r *ghttp.Request) *response.Response {
	var info model.ClusterApp
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.ClusterApp.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"clusterApp": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags ClusterApp
// @Summary 删除ClusterApp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ClusterApp模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clusterApp/delete [delete]
func (a *clusterApp) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.ClusterApp.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags ClusterApp
// @Summary 批量删除ClusterApp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除ClusterApp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clusterApp/deletes [delete]
func (a *clusterApp) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.ClusterApp.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags ClusterApp
// @Summary 分页获取ClusterApp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchClusterApp true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clusterApp/getList [get]
func (a *clusterApp) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchClusterApp
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.ClusterApp.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
