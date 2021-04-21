package model

import (
	"github.com/wenwenxiong/cloudmanager/app/model/common"
)

type Region struct {
	common.BaseModel
	ID         string `json:"id" gorm:"type:varchar(64)"`
	Name       string `json:"name" gorm:"type:varchar(256);not null;unique"`
	Datacenter string `json:"datacenter" gorm:"type:varchar(64)"`
	Provider   string `json:"provider" gorm:"type:varchar(64)"`
	Vars       string `json:"vars" gorm:"type text(65535)"`
}
