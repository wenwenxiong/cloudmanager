package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchTenant struct{
    model.Tenant
    PageInfo
}

func (s *SearchTenant) Search() g.Map {
	condition := g.Map{}
    return condition
}