package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchTenantMember struct{
    model.TenantMember
    PageInfo
}

func (s *SearchTenantMember) Search() g.Map {
	condition := g.Map{}
    return condition
}