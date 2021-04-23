package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Apptemplate = new(apptemplate)

type apptemplate struct {
	_apptemplate model.Apptemplate
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Apptemplate记录
func (s *apptemplate) Create(info *model.Apptemplate) error {
	_, err := g.DB().Table(s._apptemplate.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Apptemplate记录
func (s *apptemplate) First(info *request.GetById) (result *model.Apptemplate, err error) {
	var entity model.Apptemplate
	err = g.DB().Table(s._apptemplate.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Apptemplate记录
func (s *apptemplate) Update(info *model.Apptemplate) (result *model.Apptemplate, err error) {
    _, err = g.DB().Table(s._apptemplate.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Apptemplate记录
func (s *apptemplate) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._apptemplate.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Apptemplate记录
func (s *apptemplate) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._apptemplate.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Apptemplate记录列表
func (s *apptemplate) GetList(info *request.SearchApptemplate) (list *[]model.Apptemplate, total int, err error) {
	var apptemplates []model.Apptemplate
	db := g.DB().Table(s._apptemplate.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&apptemplates)
	return &apptemplates, total, err
}
