package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var TenantMember = new(tenantMember)

type tenantMember struct {
	_tenantMember model.TenantMember
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建TenantMember记录
func (s *tenantMember) Create(info *model.TenantMember) error {
	_, err := g.DB().Table(s._tenantMember.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取TenantMember记录
func (s *tenantMember) First(info *request.GetById) (result *model.TenantMember, err error) {
	var entity model.TenantMember
	err = g.DB().Table(s._tenantMember.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新TenantMember记录
func (s *tenantMember) Update(info *model.TenantMember) (result *model.TenantMember, err error) {
    _, err = g.DB().Table(s._tenantMember.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除TenantMember记录
func (s *tenantMember) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._tenantMember.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量TenantMember记录
func (s *tenantMember) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._tenantMember.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取TenantMember记录列表
func (s *tenantMember) GetList(info *request.SearchTenantMember) (list *[]model.TenantMember, total int, err error) {
	var tenantMembers []model.TenantMember
	db := g.DB().Table(s._tenantMember.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&tenantMembers)
	return &tenantMembers, total, err
}
