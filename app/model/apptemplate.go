package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Apptemplate struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Version  string `orm:"version" json:"version" form:"version" gorm:"column:version;comment:;type:varchar(255);size:255;"`
      Describe  string `orm:"describe" json:"describe" form:"describe" gorm:"column:describe;comment:;type:varchar(255);size:255;"`
      Status  string `orm:"status" json:"status" form:"status" gorm:"column:status;comment:;type:varchar(255);size:255;"`
      Vars  string `orm:"vars" json:"vars" form:"vars" gorm:"column:vars;comment:;type:mediumtext;"`
      Logo  string `orm:"logo" json:"logo" form:"logo" gorm:"column:logo;comment:;type:varchar(255);size:255;"`
}


func (m *Apptemplate) TableName() string {
  return "apptemplate"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ApptemplateWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Apptemplate   `json:"business"`
// }

// func (Apptemplate) TableName() string {
// 	return "apptemplate"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["apptemplate"] = func() model.GVA_Workflow {
//   return new(model.ApptemplateWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["apptemplate"] = func() interface{} {
// 	return new(model.Apptemplate)
// }