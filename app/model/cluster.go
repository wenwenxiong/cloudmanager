package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Cluster struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      SpecId  string `orm:"spec_id" json:"specId" form:"specId" gorm:"column:spec_id;comment:;type:varchar(255);size:255;"`
      SecrectId  string `orm:"secrect_id" json:"secrectId" form:"secrectId" gorm:"column:secrect_id;comment:;type:varchar(255);size:255;"`
      StatusId  string `orm:"status_id" json:"statusId" form:"statusId" gorm:"column:status_id;comment:;type:varchar(255);size:255;"`
}


func (m *Cluster) TableName() string {
  return "cluster"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClusterWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Cluster   `json:"business"`
// }

// func (Cluster) TableName() string {
// 	return "cluster"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["cluster"] = func() model.GVA_Workflow {
//   return new(model.ClusterWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["cluster"] = func() interface{} {
// 	return new(model.Cluster)
// }