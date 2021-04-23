package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchTenantResource struct{
    model.TenantResource
    PageInfo
}

func (s *SearchTenantResource) Search() g.Map {
	condition := g.Map{}
    return condition
}