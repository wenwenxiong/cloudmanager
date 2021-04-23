package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchClusterNodes struct{
    model.ClusterNodes
    PageInfo
}

func (s *SearchClusterNodes) Search() g.Map {
	condition := g.Map{}
    return condition
}