package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchSystemSetting struct{
    model.SystemSetting
    PageInfo
}

func (s *SearchSystemSetting) Search() g.Map {
	condition := g.Map{}
    return condition
}