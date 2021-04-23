package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type UserClusterResourceLimit struct {
      global.Model
      UserId  string `orm:"user_id" json:"userId" form:"userId" gorm:"column:user_id;comment:;type:varchar(64);size:64;"`
      ClusterId  string `orm:"cluster_id" json:"clusterId" form:"clusterId" gorm:"column:cluster_id;comment:;type:varchar(255);size:255;"`
      MaxCpuCore  int `orm:"max_cpu_core" json:"maxCpuCore" form:"maxCpuCore" gorm:"column:max_cpu_core;comment:;type:int;size:10;"`
      MaxMemory  int `orm:"max_memory" json:"maxMemory" form:"maxMemory" gorm:"column:max_memory;comment:;type:int;size:10;"`
      MaxPod  int `orm:"max_pod" json:"maxPod" form:"maxPod" gorm:"column:max_pod;comment:;type:int;size:10;"`
}


func (m *UserClusterResourceLimit) TableName() string {
  return "user_cluster_resource_limit"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type UserClusterResourceLimitWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	UserClusterResourceLimit   `json:"business"`
// }

// func (UserClusterResourceLimit) TableName() string {
// 	return "user_cluster_resource_limit"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["userClusterResourceLimit"] = func() model.GVA_Workflow {
//   return new(model.UserClusterResourceLimitWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["userClusterResourceLimit"] = func() interface{} {
// 	return new(model.UserClusterResourceLimit)
// }