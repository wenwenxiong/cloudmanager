package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Tenant = new(tenant)

type tenant struct {
	_tenant model.Tenant
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Tenant记录
func (s *tenant) Create(info *model.Tenant) error {
	_, err := g.DB().Table(s._tenant.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Tenant记录
func (s *tenant) First(info *request.GetById) (result *model.Tenant, err error) {
	var entity model.Tenant
	err = g.DB().Table(s._tenant.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Tenant记录
func (s *tenant) Update(info *model.Tenant) (result *model.Tenant, err error) {
    _, err = g.DB().Table(s._tenant.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Tenant记录
func (s *tenant) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._tenant.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Tenant记录
func (s *tenant) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._tenant.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Tenant记录列表
func (s *tenant) GetList(info *request.SearchTenant) (list *[]model.Tenant, total int, err error) {
	var tenants []model.Tenant
	db := g.DB().Table(s._tenant.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&tenants)
	return &tenants, total, err
}
