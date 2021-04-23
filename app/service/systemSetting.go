package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var SystemSetting = new(systemSetting)

type systemSetting struct {
	_systemSetting model.SystemSetting
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建SystemSetting记录
func (s *systemSetting) Create(info *model.SystemSetting) error {
	_, err := g.DB().Table(s._systemSetting.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取SystemSetting记录
func (s *systemSetting) First(info *request.GetById) (result *model.SystemSetting, err error) {
	var entity model.SystemSetting
	err = g.DB().Table(s._systemSetting.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新SystemSetting记录
func (s *systemSetting) Update(info *model.SystemSetting) (result *model.SystemSetting, err error) {
    _, err = g.DB().Table(s._systemSetting.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除SystemSetting记录
func (s *systemSetting) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._systemSetting.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量SystemSetting记录
func (s *systemSetting) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._systemSetting.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取SystemSetting记录列表
func (s *systemSetting) GetList(info *request.SearchSystemSetting) (list *[]model.SystemSetting, total int, err error) {
	var systemSettings []model.SystemSetting
	db := g.DB().Table(s._systemSetting.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&systemSettings)
	return &systemSettings, total, err
}
