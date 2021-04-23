package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type Host struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Ip  string `orm:"ip" json:"ip" form:"ip" gorm:"column:ip;comment:;type:varchar(128);size:128;"`
      Os  string `orm:"os" json:"os" form:"os" gorm:"column:os;comment:;type:varchar(64);size:64;"`
      OsVersion  string `orm:"os_version" json:"osVersion" form:"osVersion" gorm:"column:os_version;comment:;type:varchar(64);size:64;"`
      CpuCore  int `orm:"cpu_core" json:"cpuCore" form:"cpuCore" gorm:"column:cpu_core;comment:;type:int;size:10;"`
      Memory  int `orm:"memory" json:"memory" form:"memory" gorm:"column:memory;comment:;type:int;size:10;"`
      HasGpu  *bool `orm:"has_gpu" json:"hasGpu" form:"hasGpu" gorm:"column:has_gpu;comment:;type:tinyint;"`
      GpuNum  int `orm:"gpu_num" json:"gpuNum" form:"gpuNum" gorm:"column:gpu_num;comment:;type:int;size:10;"`
      GpuInfo  string `orm:"gpu_info" json:"gpuInfo" form:"gpuInfo" gorm:"column:gpu_info;comment:;type:varchar(128);size:128;"`
      Port  string `orm:"port" json:"port" form:"port" gorm:"column:port;comment:;type:varchar(64);size:64;"`
      Status  string `orm:"status" json:"status" form:"status" gorm:"column:status;comment:;type:varchar(64);size:64;"`
      Tags  string `orm:"tags" json:"tags" form:"tags" gorm:"column:tags;comment:;type:varchar(255);size:255;"`
      CredentialId  string `orm:"credential_id" json:"credentialId" form:"credentialId" gorm:"column:credential_id;comment:;type:varchar(64);size:64;"`
      ZoneId  string `orm:"zone_id" json:"zoneId" form:"zoneId" gorm:"column:zone_id;comment:;type:varchar(255);size:255;"`
      Message  string `orm:"message" json:"message" form:"message" gorm:"column:message;comment:;type:mediumtext;"`
      Architectures  string `orm:"architectures" json:"architectures" form:"architectures" gorm:"column:architectures;comment:;type:varchar(255);size:255;"`
}


func (m *Host) TableName() string {
  return "host"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type HostWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Host   `json:"business"`
// }

// func (Host) TableName() string {
// 	return "host"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["host"] = func() model.GVA_Workflow {
//   return new(model.HostWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["host"] = func() interface{} {
// 	return new(model.Host)
// }