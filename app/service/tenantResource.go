package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var TenantResource = new(tenantResource)

type tenantResource struct {
	_tenantResource model.TenantResource
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建TenantResource记录
func (s *tenantResource) Create(info *model.TenantResource) error {
	_, err := g.DB().Table(s._tenantResource.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取TenantResource记录
func (s *tenantResource) First(info *request.GetById) (result *model.TenantResource, err error) {
	var entity model.TenantResource
	err = g.DB().Table(s._tenantResource.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新TenantResource记录
func (s *tenantResource) Update(info *model.TenantResource) (result *model.TenantResource, err error) {
    _, err = g.DB().Table(s._tenantResource.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除TenantResource记录
func (s *tenantResource) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._tenantResource.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量TenantResource记录
func (s *tenantResource) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._tenantResource.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取TenantResource记录列表
func (s *tenantResource) GetList(info *request.SearchTenantResource) (list *[]model.TenantResource, total int, err error) {
	var tenantResources []model.TenantResource
	db := g.DB().Table(s._tenantResource.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&tenantResources)
	return &tenantResources, total, err
}
