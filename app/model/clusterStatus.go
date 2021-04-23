package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type ClusterStatus struct {
      global.Model
      Message  string `orm:"message" json:"message" form:"message" gorm:"column:message;comment:;type:mediumtext;"`
      Phase  string `orm:"phase" json:"phase" form:"phase" gorm:"column:phase;comment:;type:varchar(255);size:255;"`
}


func (m *ClusterStatus) TableName() string {
  return "cluster_status"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClusterStatusWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	ClusterStatus   `json:"business"`
// }

// func (ClusterStatus) TableName() string {
// 	return "cluster_status"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["clusterStatus"] = func() model.GVA_Workflow {
//   return new(model.ClusterStatusWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["clusterStatus"] = func() interface{} {
// 	return new(model.ClusterStatus)
// }