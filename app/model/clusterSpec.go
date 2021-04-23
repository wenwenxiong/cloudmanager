package model

import "cloudmanager/library/global"

// 如果含有time.Time 请自行import time包
type ClusterSpec struct {
      global.Model
      Version  string `orm:"version" json:"version" form:"version" gorm:"column:version;comment:;type:varchar(255);size:255;"`
      NetworkType  string `orm:"network_type" json:"networkType" form:"networkType" gorm:"column:network_type;comment:;type:varchar(255);size:255;"`
      RuntimeType  string `orm:"runtime_type" json:"runtimeType" form:"runtimeType" gorm:"column:runtime_type;comment:;type:varchar(255);size:255;"`
      DockerStorageDir  string `orm:"docker_storage_dir" json:"dockerStorageDir" form:"dockerStorageDir" gorm:"column:docker_storage_dir;comment:;type:varchar(255);size:255;"`
      ContainerdStorageDir  string `orm:"containerd_storage_dir" json:"containerdStorageDir" form:"containerdStorageDir" gorm:"column:containerd_storage_dir;comment:;type:varchar(255);size:255;"`
      LbKubeApiserverIp  string `orm:"lb_kube_apiserver_ip" json:"lbKubeApiserverIp" form:"lbKubeApiserverIp" gorm:"column:lb_kube_apiserver_ip;comment:;type:varchar(255);size:255;"`
      KubeApiServerPort  string `orm:"kube_api_server_port" json:"kubeApiServerPort" form:"kubeApiServerPort" gorm:"column:kube_api_server_port;comment:;type:varchar(255);size:255;"`
      KubePodSubnet  string `orm:"kube_pod_subnet" json:"kubePodSubnet" form:"kubePodSubnet" gorm:"column:kube_pod_subnet;comment:;type:varchar(255);size:255;"`
      KubeServiceSubnet  string `orm:"kube_service_subnet" json:"kubeServiceSubnet" form:"kubeServiceSubnet" gorm:"column:kube_service_subnet;comment:;type:varchar(255);size:255;"`
      WorkerAmount  int `orm:"worker_amount" json:"workerAmount" form:"workerAmount" gorm:"column:worker_amount;comment:;type:int;size:10;"`
      KubeMaxPods  int `orm:"kube_max_pods" json:"kubeMaxPods" form:"kubeMaxPods" gorm:"column:kube_max_pods;comment:;type:int;size:10;"`
      KubeProxyMode  string `orm:"kube_proxy_mode" json:"kubeProxyMode" form:"kubeProxyMode" gorm:"column:kube_proxy_mode;comment:;type:varchar(255);size:255;"`
      Architectures  string `orm:"architectures" json:"architectures" form:"architectures" gorm:"column:architectures;comment:;type:varchar(255);size:255;"`
      SupportGpu  string `orm:"support_gpu" json:"supportGpu" form:"supportGpu" gorm:"column:support_gpu;comment:;type:varchar(255);size:255;"`
}


func (m *ClusterSpec) TableName() string {
  return "cluster_spec"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClusterSpecWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	ClusterSpec   `json:"business"`
// }

// func (ClusterSpec) TableName() string {
// 	return "cluster_spec"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["clusterSpec"] = func() model.GVA_Workflow {
//   return new(model.ClusterSpecWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["clusterSpec"] = func() interface{} {
// 	return new(model.ClusterSpec)
// }