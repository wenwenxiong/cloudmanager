package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Zone struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Vars  string `orm:"vars" json:"vars" form:"vars" gorm:"column:vars;comment:;type:mediumtext;"`
      Status  string `orm:"status" json:"status" form:"status" gorm:"column:status;comment:;type:varchar(64);size:64;"`
      RegionId  string `orm:"region_id" json:"regionId" form:"regionId" gorm:"column:region_id;comment:;type:varchar(64);size:64;"`
}


func (m *Zone) TableName() string {
  return "zone"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ZoneWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Zone   `json:"business"`
// }

// func (Zone) TableName() string {
// 	return "zone"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["zone"] = func() model.GVA_Workflow {
//   return new(model.ZoneWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["zone"] = func() interface{} {
// 	return new(model.Zone)
// }