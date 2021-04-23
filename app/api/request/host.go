package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchHost struct{
    model.Host
    PageInfo
}

func (s *SearchHost) Search() g.Map {
	condition := g.Map{}
    return condition
}