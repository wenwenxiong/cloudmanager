package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Region = new(region)

type region struct {
	_region model.Region
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Region记录
func (s *region) Create(info *model.Region) error {
	_, err := g.DB().Table(s._region.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Region记录
func (s *region) First(info *request.GetById) (result *model.Region, err error) {
	var entity model.Region
	err = g.DB().Table(s._region.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Region记录
func (s *region) Update(info *model.Region) (result *model.Region, err error) {
    _, err = g.DB().Table(s._region.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Region记录
func (s *region) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._region.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Region记录
func (s *region) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._region.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Region记录列表
func (s *region) GetList(info *request.SearchRegion) (list *[]model.Region, total int, err error) {
	var regions []model.Region
	db := g.DB().Table(s._region.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&regions)
	return &regions, total, err
}
