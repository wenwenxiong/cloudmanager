package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ClusterNodes = new(clusterNodes)

type clusterNodes struct{}

// @Tags ClusterNodes
// @Summary 创建ClusterNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterNodes true "创建ClusterNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clusterNodes/create [post]
func (a *clusterNodes) Create(r *ghttp.Request) *response.Response {
	var info model.ClusterNodes
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.ClusterNodes.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags ClusterNodes
// @Summary 用id查询ClusterNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clusterNodes/first [get]
func (a *clusterNodes) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.ClusterNodes.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"clusterNodes": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags ClusterNodes
// @Summary 更新ClusterNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterNodes true "clusterNodes模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clusterNodes/update [put]
func (a *clusterNodes) Update(r *ghttp.Request) *response.Response {
	var info model.ClusterNodes
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.ClusterNodes.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"clusterNodes": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags ClusterNodes
// @Summary 删除ClusterNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ClusterNodes模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clusterNodes/delete [delete]
func (a *clusterNodes) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.ClusterNodes.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags ClusterNodes
// @Summary 批量删除ClusterNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除ClusterNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clusterNodes/deletes [delete]
func (a *clusterNodes) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.ClusterNodes.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags ClusterNodes
// @Summary 分页获取ClusterNodes列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchClusterNodes true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clusterNodes/getList [get]
func (a *clusterNodes) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchClusterNodes
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.ClusterNodes.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
