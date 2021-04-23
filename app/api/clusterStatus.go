package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ClusterStatus = new(clusterStatus)

type clusterStatus struct{}

// @Tags ClusterStatus
// @Summary 创建ClusterStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterStatus true "创建ClusterStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clusterStatus/create [post]
func (a *clusterStatus) Create(r *ghttp.Request) *response.Response {
	var info model.ClusterStatus
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.ClusterStatus.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags ClusterStatus
// @Summary 用id查询ClusterStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clusterStatus/first [get]
func (a *clusterStatus) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.ClusterStatus.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"clusterStatus": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags ClusterStatus
// @Summary 更新ClusterStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterStatus true "clusterStatus模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clusterStatus/update [put]
func (a *clusterStatus) Update(r *ghttp.Request) *response.Response {
	var info model.ClusterStatus
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.ClusterStatus.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"clusterStatus": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags ClusterStatus
// @Summary 删除ClusterStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ClusterStatus模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clusterStatus/delete [delete]
func (a *clusterStatus) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.ClusterStatus.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags ClusterStatus
// @Summary 批量删除ClusterStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除ClusterStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clusterStatus/deletes [delete]
func (a *clusterStatus) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.ClusterStatus.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags ClusterStatus
// @Summary 分页获取ClusterStatus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchClusterStatus true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clusterStatus/getList [get]
func (a *clusterStatus) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchClusterStatus
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.ClusterStatus.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
