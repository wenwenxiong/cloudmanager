package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchVolume struct{
    model.Volume
    PageInfo
}

func (s *SearchVolume) Search() g.Map {
	condition := g.Map{}
    return condition
}