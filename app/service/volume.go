package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Volume = new(volume)

type volume struct {
	_volume model.Volume
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Volume记录
func (s *volume) Create(info *model.Volume) error {
	_, err := g.DB().Table(s._volume.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Volume记录
func (s *volume) First(info *request.GetById) (result *model.Volume, err error) {
	var entity model.Volume
	err = g.DB().Table(s._volume.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Volume记录
func (s *volume) Update(info *model.Volume) (result *model.Volume, err error) {
    _, err = g.DB().Table(s._volume.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Volume记录
func (s *volume) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._volume.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Volume记录
func (s *volume) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._volume.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Volume记录列表
func (s *volume) GetList(info *request.SearchVolume) (list *[]model.Volume, total int, err error) {
	var volumes []model.Volume
	db := g.DB().Table(s._volume.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&volumes)
	return &volumes, total, err
}
