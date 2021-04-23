package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchClusterStatus struct{
    model.ClusterStatus
    PageInfo
}

func (s *SearchClusterStatus) Search() g.Map {
	condition := g.Map{}
    return condition
}