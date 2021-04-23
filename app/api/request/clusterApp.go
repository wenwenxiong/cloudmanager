package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchClusterApp struct{
    model.ClusterApp
    PageInfo
}

func (s *SearchClusterApp) Search() g.Map {
	condition := g.Map{}
    return condition
}