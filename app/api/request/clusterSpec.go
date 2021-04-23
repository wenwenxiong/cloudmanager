package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchClusterSpec struct{
    model.ClusterSpec
    PageInfo
}

func (s *SearchClusterSpec) Search() g.Map {
	condition := g.Map{}
    return condition
}