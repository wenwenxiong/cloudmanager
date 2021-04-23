package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchRegion struct{
    model.Region
    PageInfo
}

func (s *SearchRegion) Search() g.Map {
	condition := g.Map{}
    return condition
}