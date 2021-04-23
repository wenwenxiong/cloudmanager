package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchCredential struct{
    model.Credential
    PageInfo
}

func (s *SearchCredential) Search() g.Map {
	condition := g.Map{}
    return condition
}