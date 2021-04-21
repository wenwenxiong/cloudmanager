package model

import "github.com/wenwenxiong/cloudmanager/app/model/common"

type Host struct {
	common.BaseModel
	ID           string     `json:"_"`
	Name         string     `json:"name" gorm:"type:varchar(256);not null;unique"`
	Memory       int        `json:"memory" gorm:"type:int(64)"`
	CpuCore      int        `json:"cpuCore" gorm:"type:int(64)"`
	Os           string     `json:"os" gorm:"type:varchar(64)"`
	OsVersion    string     `json:"osVersion" gorm:"type:varchar(64)"`
	GpuNum       int        `json:"gpuNum" gorm:"type:int(64)"`
	GpuInfo      string     `json:"gpuInfo" gorm:"type:varchar(128)"`
	Ip           string     `json:"ip" gorm:"type:varchar(128);not null;unique"`
	Port         int        `json:"port" gorm:"type:varchar(64)"`
	CredentialID string     `json:"credentialId" gorm:"type:varchar(64)"`
	ClusterID    string     `json:"clusterId" gorm:"type:varchar(64)"`
	ZoneID       string     `json:"zoneId" gorm:"type:varchar(64)"`
	Zone         Zone       `json:"-"  gorm:"save_associations:false" `
	Volumes      []Volume   `json:"volumes" gorm:"save_associations:false"`
	Credential   Credential `json:"-" gorm:"save_associations:false" `
	Status       string     `json:"status" gorm:"type:varchar(64)"`
	Message      string     `json:"message" gorm:"type:text(65535)"`
}
