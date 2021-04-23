package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type SystemSetting struct {
      global.Model
      Key  string `orm:"key" json:"key" form:"key" gorm:"column:key;comment:;type:varchar(255);size:255;"`
      Value  string `orm:"value" json:"value" form:"value" gorm:"column:value;comment:;type:varchar(255);size:255;"`
      Tab  string `orm:"tab" json:"tab" form:"tab" gorm:"column:tab;comment:;type:varchar(64);size:64;"`
}


func (m *SystemSetting) TableName() string {
  return "system_setting"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type SystemSettingWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	SystemSetting   `json:"business"`
// }

// func (SystemSetting) TableName() string {
// 	return "system_setting"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["systemSetting"] = func() model.GVA_Workflow {
//   return new(model.SystemSettingWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["systemSetting"] = func() interface{} {
// 	return new(model.SystemSetting)
// }