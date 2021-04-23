package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Region struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Datecenter  string `orm:"datecenter" json:"datecenter" form:"datecenter" gorm:"column:datecenter;comment:;type:varchar(64);size:64;"`
      Provider  string `orm:"provider" json:"provider" form:"provider" gorm:"column:provider;comment:;type:varchar(64);size:64;"`
      Vars  string `orm:"vars" json:"vars" form:"vars" gorm:"column:vars;comment:;type:mediumtext;"`
}


func (m *Region) TableName() string {
  return "region"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type RegionWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Region   `json:"business"`
// }

// func (Region) TableName() string {
// 	return "region"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["region"] = func() model.GVA_Workflow {
//   return new(model.RegionWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["region"] = func() interface{} {
// 	return new(model.Region)
// }