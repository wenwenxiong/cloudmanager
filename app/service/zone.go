package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Zone = new(zone)

type zone struct {
	_zone model.Zone
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Zone记录
func (s *zone) Create(info *model.Zone) error {
	_, err := g.DB().Table(s._zone.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Zone记录
func (s *zone) First(info *request.GetById) (result *model.Zone, err error) {
	var entity model.Zone
	err = g.DB().Table(s._zone.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Zone记录
func (s *zone) Update(info *model.Zone) (result *model.Zone, err error) {
    _, err = g.DB().Table(s._zone.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Zone记录
func (s *zone) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._zone.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Zone记录
func (s *zone) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._zone.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Zone记录列表
func (s *zone) GetList(info *request.SearchZone) (list *[]model.Zone, total int, err error) {
	var zones []model.Zone
	db := g.DB().Table(s._zone.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&zones)
	return &zones, total, err
}
