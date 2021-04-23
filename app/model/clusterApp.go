package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type ClusterApp struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      ClusterId  string `orm:"cluster_id" json:"clusterId" form:"clusterId" gorm:"column:cluster_id;comment:;type:varchar(255);size:255;"`
      TenantId  string `orm:"tenant_id" json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:;type:varchar(255);size:255;"`
      ApptemplateId  string `orm:"apptemplate_id" json:"apptemplateId" form:"apptemplateId" gorm:"column:apptemplate_id;comment:;type:varchar(64);size:64;"`
      Status  string `orm:"status" json:"status" form:"status" gorm:"column:status;comment:;type:varchar(255);size:255;"`
      Vars  string `orm:"vars" json:"vars" form:"vars" gorm:"column:vars;comment:;type:mediumtext;"`
}


func (m *ClusterApp) TableName() string {
  return "cluster_app"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClusterAppWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	ClusterApp   `json:"business"`
// }

// func (ClusterApp) TableName() string {
// 	return "cluster_app"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["clusterApp"] = func() model.GVA_Workflow {
//   return new(model.ClusterAppWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["clusterApp"] = func() interface{} {
// 	return new(model.ClusterApp)
// }