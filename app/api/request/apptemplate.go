package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchApptemplate struct{
    model.Apptemplate
    PageInfo
}

func (s *SearchApptemplate) Search() g.Map {
	condition := g.Map{}
    return condition
}