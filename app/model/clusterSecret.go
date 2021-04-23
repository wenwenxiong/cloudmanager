package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type ClusterSecret struct {
      global.Model
      KubeadmToken  string `orm:"kubeadm_token" json:"kubeadmToken" form:"kubeadmToken" gorm:"column:kubeadm_token;comment:;type:mediumtext;"`
      KubernetesToken  string `orm:"kubernetes_token" json:"kubernetesToken" form:"kubernetesToken" gorm:"column:kubernetes_token;comment:;type:mediumtext;"`
}


func (m *ClusterSecret) TableName() string {
  return "cluster_secret"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClusterSecretWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	ClusterSecret   `json:"business"`
// }

// func (ClusterSecret) TableName() string {
// 	return "cluster_secret"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["clusterSecret"] = func() model.GVA_Workflow {
//   return new(model.ClusterSecretWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["clusterSecret"] = func() interface{} {
// 	return new(model.ClusterSecret)
// }