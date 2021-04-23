package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type User struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Password  string `orm:"password" json:"password" form:"password" gorm:"column:password;comment:;type:varchar(255);size:255;"`
      Email  string `orm:"email" json:"email" form:"email" gorm:"column:email;comment:;type:varchar(255);size:255;"`
      IsActive  *bool `orm:"is_active" json:"isActive" form:"isActive" gorm:"column:is_active;comment:;type:tinyint;"`
      Language  string `orm:"language" json:"language" form:"language" gorm:"column:language;comment:;type:varchar(64);size:64;"`
      IsAdmin  *bool `orm:"is_admin" json:"isAdmin" form:"isAdmin" gorm:"column:is_admin;comment:;type:tinyint;"`
}


func (m *User) TableName() string {
  return "user"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type UserWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	User   `json:"business"`
// }

// func (User) TableName() string {
// 	return "user"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["user"] = func() model.GVA_Workflow {
//   return new(model.UserWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["user"] = func() interface{} {
// 	return new(model.User)
// }