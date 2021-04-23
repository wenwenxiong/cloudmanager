package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchClusterSecret struct{
    model.ClusterSecret
    PageInfo
}

func (s *SearchClusterSecret) Search() g.Map {
	condition := g.Map{}
    return condition
}