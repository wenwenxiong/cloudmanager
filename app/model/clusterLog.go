package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type ClusterLog struct {
      global.Model
      ClusterId  string `orm:"cluster_id" json:"clusterId" form:"clusterId" gorm:"column:cluster_id;comment:;type:varchar(255);size:255;"`
      Type  string `orm:"type" json:"type" form:"type" gorm:"column:type;comment:;type:varchar(255);size:255;"`
      Message  string `orm:"message" json:"message" form:"message" gorm:"column:message;comment:;type:mediumtext;"`
}


func (m *ClusterLog) TableName() string {
  return "cluster_log"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClusterLogWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	ClusterLog   `json:"business"`
// }

// func (ClusterLog) TableName() string {
// 	return "cluster_log"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["clusterLog"] = func() model.GVA_Workflow {
//   return new(model.ClusterLogWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["clusterLog"] = func() interface{} {
// 	return new(model.ClusterLog)
// }