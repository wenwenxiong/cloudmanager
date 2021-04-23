package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Credential struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Username  string `orm:"username" json:"username" form:"username" gorm:"column:username;comment:;type:varchar(255);size:255;"`
      Password  string `orm:"password" json:"password" form:"password" gorm:"column:password;comment:;type:varchar(255);size:255;"`
      PrivateKey  string `orm:"private_key" json:"privateKey" form:"privateKey" gorm:"column:private_key;comment:;type:mediumtext;"`
      Type  string `orm:"type" json:"type" form:"type" gorm:"column:type;comment:;type:varchar(255);size:255;"`
}


func (m *Credential) TableName() string {
  return "credential"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type CredentialWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Credential   `json:"business"`
// }

// func (Credential) TableName() string {
// 	return "credential"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["credential"] = func() model.GVA_Workflow {
//   return new(model.CredentialWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["credential"] = func() interface{} {
// 	return new(model.Credential)
// }