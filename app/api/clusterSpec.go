package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ClusterSpec = new(clusterSpec)

type clusterSpec struct{}

// @Tags ClusterSpec
// @Summary 创建ClusterSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterSpec true "创建ClusterSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clusterSpec/create [post]
func (a *clusterSpec) Create(r *ghttp.Request) *response.Response {
	var info model.ClusterSpec
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.ClusterSpec.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags ClusterSpec
// @Summary 用id查询ClusterSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clusterSpec/first [get]
func (a *clusterSpec) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.ClusterSpec.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"clusterSpec": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags ClusterSpec
// @Summary 更新ClusterSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterSpec true "clusterSpec模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clusterSpec/update [put]
func (a *clusterSpec) Update(r *ghttp.Request) *response.Response {
	var info model.ClusterSpec
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.ClusterSpec.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"clusterSpec": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags ClusterSpec
// @Summary 删除ClusterSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ClusterSpec模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clusterSpec/delete [delete]
func (a *clusterSpec) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.ClusterSpec.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags ClusterSpec
// @Summary 批量删除ClusterSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除ClusterSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clusterSpec/deletes [delete]
func (a *clusterSpec) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.ClusterSpec.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags ClusterSpec
// @Summary 分页获取ClusterSpec列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchClusterSpec true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clusterSpec/getList [get]
func (a *clusterSpec) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchClusterSpec
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.ClusterSpec.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
