package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Credential = new(credential)

type credential struct {
	_credential model.Credential
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Credential记录
func (s *credential) Create(info *model.Credential) error {
	_, err := g.DB().Table(s._credential.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Credential记录
func (s *credential) First(info *request.GetById) (result *model.Credential, err error) {
	var entity model.Credential
	err = g.DB().Table(s._credential.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Credential记录
func (s *credential) Update(info *model.Credential) (result *model.Credential, err error) {
    _, err = g.DB().Table(s._credential.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Credential记录
func (s *credential) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._credential.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Credential记录
func (s *credential) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._credential.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Credential记录列表
func (s *credential) GetList(info *request.SearchCredential) (list *[]model.Credential, total int, err error) {
	var credentials []model.Credential
	db := g.DB().Table(s._credential.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&credentials)
	return &credentials, total, err
}
