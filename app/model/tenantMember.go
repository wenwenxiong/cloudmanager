package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type TenantMember struct {
      global.Model
      TenantId  string `orm:"tenant_id" json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:;type:varchar(64);size:64;"`
      UserId  string `orm:"user_id" json:"userId" form:"userId" gorm:"column:user_id;comment:;type:varchar(64);size:64;"`
      Role  string `orm:"role" json:"role" form:"role" gorm:"column:role;comment:;type:varchar(64);size:64;"`
}


func (m *TenantMember) TableName() string {
  return "tenant_member"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type TenantMemberWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	TenantMember   `json:"business"`
// }

// func (TenantMember) TableName() string {
// 	return "tenant_member"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["tenantMember"] = func() model.GVA_Workflow {
//   return new(model.TenantMemberWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["tenantMember"] = func() interface{} {
// 	return new(model.TenantMember)
// }