package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type ClusterNodes struct {
      global.Model
      Name  string `orm:"name" json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      HostId  string `orm:"host_id" json:"hostId" form:"hostId" gorm:"column:host_id;comment:;type:varchar(255);size:255;"`
      ClusterId  string `orm:"cluster_id" json:"clusterId" form:"clusterId" gorm:"column:cluster_id;comment:;type:varchar(255);size:255;"`
      Role  string `orm:"role" json:"role" form:"role" gorm:"column:role;comment:;type:varchar(255);size:255;"`
      Status  string `orm:"status" json:"status" form:"status" gorm:"column:status;comment:;type:varchar(255);size:255;"`
      Message  string `orm:"message" json:"message" form:"message" gorm:"column:message;comment:;type:mediumtext;"`
}


func (m *ClusterNodes) TableName() string {
  return "cluster_nodes"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClusterNodesWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	ClusterNodes   `json:"business"`
// }

// func (ClusterNodes) TableName() string {
// 	return "cluster_nodes"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["clusterNodes"] = func() model.GVA_Workflow {
//   return new(model.ClusterNodesWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["clusterNodes"] = func() interface{} {
// 	return new(model.ClusterNodes)
// }