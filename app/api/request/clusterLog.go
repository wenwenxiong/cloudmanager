package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchClusterLog struct{
    model.ClusterLog
    PageInfo
}

func (s *SearchClusterLog) Search() g.Map {
	condition := g.Map{}
    return condition
}