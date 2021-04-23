package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type TenantResource struct {
      global.Model
      ResourceType  string `orm:"resource_type" json:"resourceType" form:"resourceType" gorm:"column:resource_type;comment:;type:varchar(128);size:128;"`
      ResourceId  string `orm:"resource_id" json:"resourceId" form:"resourceId" gorm:"column:resource_id;comment:;type:varchar(64);size:64;"`
      TenantId  string `orm:"tenant_id" json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:;type:varchar(64);size:64;"`
}


func (m *TenantResource) TableName() string {
  return "tenant_resource"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type TenantResourceWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	TenantResource   `json:"business"`
// }

// func (TenantResource) TableName() string {
// 	return "tenant_resource"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["tenantResource"] = func() model.GVA_Workflow {
//   return new(model.TenantResourceWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["tenantResource"] = func() interface{} {
// 	return new(model.TenantResource)
// }