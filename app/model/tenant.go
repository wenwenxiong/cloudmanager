package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Tenant struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Description  string `orm:"description" json:"description" form:"description" gorm:"column:description;comment:;type:varchar(128);size:128;"`
}


func (m *Tenant) TableName() string {
  return "tenant"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type TenantWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Tenant   `json:"business"`
// }

// func (Tenant) TableName() string {
// 	return "tenant"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["tenant"] = func() model.GVA_Workflow {
//   return new(model.TenantWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["tenant"] = func() interface{} {
// 	return new(model.Tenant)
// }