package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Volume struct {
      global.Model
      HostId  string `orm:"host_id" json:"hostId" form:"hostId" gorm:"column:host_id;comment:;type:varchar(64);size:64;"`
      Size  int `orm:"size" json:"size" form:"size" gorm:"column:size;comment:;type:int;size:10;"`
      Name  int `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:int;size:10;"`
}


func (m *Volume) TableName() string {
  return "volume"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type VolumeWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Volume   `json:"business"`
// }

// func (Volume) TableName() string {
// 	return "volume"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["volume"] = func() model.GVA_Workflow {
//   return new(model.VolumeWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["volume"] = func() interface{} {
// 	return new(model.Volume)
// }