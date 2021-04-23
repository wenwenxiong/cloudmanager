package request

import (
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

type SearchUser struct{
    model.User
    PageInfo
}

func (s *SearchUser) Search() g.Map {
	condition := g.Map{}
    return condition
}