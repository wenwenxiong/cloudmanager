package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchUserClusterResourceLimit struct{
    model.UserClusterResourceLimit
    PageInfo
}

func (s *SearchUserClusterResourceLimit) Search() g.Map {
	condition := g.Map{}
    return condition
}