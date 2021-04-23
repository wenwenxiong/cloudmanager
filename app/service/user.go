package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var User = new(user)

type user struct {
	_user model.User
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建User记录
func (s *user) Create(info *model.User) error {
	_, err := g.DB().Table(s._user.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取User记录
func (s *user) First(info *request.GetById) (result *model.User, err error) {
	var entity model.User
	err = g.DB().Table(s._user.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新User记录
func (s *user) Update(info *model.User) (result *model.User, err error) {
    _, err = g.DB().Table(s._user.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除User记录
func (s *user) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._user.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量User记录
func (s *user) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._user.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取User记录列表
func (s *user) GetList(info *request.SearchUser) (list *[]model.User, total int, err error) {
	var users []model.User
	db := g.DB().Table(s._user.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&users)
	return &users, total, err
}
