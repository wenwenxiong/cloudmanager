package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchZone struct{
    model.Zone
    PageInfo
}

func (s *SearchZone) Search() g.Map {
	condition := g.Map{}
    return condition
}