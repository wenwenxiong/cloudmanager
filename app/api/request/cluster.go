package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchCluster struct{
    model.Cluster
    PageInfo
}

func (s *SearchCluster) Search() g.Map {
	condition := g.Map{}
    return condition
}