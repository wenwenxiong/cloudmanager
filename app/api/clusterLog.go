package api

import (
	"cloudmanager/app/api/request"
	"cloudmanager/library/response"
	"cloudmanager/app/model"
	"cloudmanager/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ClusterLog = new(clusterLog)

type clusterLog struct{}

// @Tags ClusterLog
// @Summary 创建ClusterLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterLog true "创建ClusterLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clusterLog/create [post]
func (a *clusterLog) Create(r *ghttp.Request) *response.Response {
	var info model.ClusterLog
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.ClusterLog.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags ClusterLog
// @Summary 用id查询ClusterLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clusterLog/first [get]
func (a *clusterLog) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.ClusterLog.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"clusterLog": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags ClusterLog
// @Summary 更新ClusterLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClusterLog true "clusterLog模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clusterLog/update [put]
func (a *clusterLog) Update(r *ghttp.Request) *response.Response {
	var info model.ClusterLog
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.ClusterLog.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"clusterLog": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags ClusterLog
// @Summary 删除ClusterLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ClusterLog模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clusterLog/delete [delete]
func (a *clusterLog) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.ClusterLog.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags ClusterLog
// @Summary 批量删除ClusterLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除ClusterLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clusterLog/deletes [delete]
func (a *clusterLog) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.ClusterLog.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags ClusterLog
// @Summary 分页获取ClusterLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchClusterLog true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clusterLog/getList [get]
func (a *clusterLog) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchClusterLog
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.ClusterLog.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
