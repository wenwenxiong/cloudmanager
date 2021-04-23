package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var UserClusterResourceLimit = new(userClusterResourceLimit)

type userClusterResourceLimit struct {
	_userClusterResourceLimit model.UserClusterResourceLimit
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建UserClusterResourceLimit记录
func (s *userClusterResourceLimit) Create(info *model.UserClusterResourceLimit) error {
	_, err := g.DB().Table(s._userClusterResourceLimit.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取UserClusterResourceLimit记录
func (s *userClusterResourceLimit) First(info *request.GetById) (result *model.UserClusterResourceLimit, err error) {
	var entity model.UserClusterResourceLimit
	err = g.DB().Table(s._userClusterResourceLimit.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新UserClusterResourceLimit记录
func (s *userClusterResourceLimit) Update(info *model.UserClusterResourceLimit) (result *model.UserClusterResourceLimit, err error) {
    _, err = g.DB().Table(s._userClusterResourceLimit.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除UserClusterResourceLimit记录
func (s *userClusterResourceLimit) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._userClusterResourceLimit.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量UserClusterResourceLimit记录
func (s *userClusterResourceLimit) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._userClusterResourceLimit.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取UserClusterResourceLimit记录列表
func (s *userClusterResourceLimit) GetList(info *request.SearchUserClusterResourceLimit) (list *[]model.UserClusterResourceLimit, total int, err error) {
	var userClusterResourceLimits []model.UserClusterResourceLimit
	db := g.DB().Table(s._userClusterResourceLimit.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&userClusterResourceLimits)
	return &userClusterResourceLimits, total, err
}
