package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Host = new(host)

type host struct {
	_host model.Host
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Host记录
func (s *host) Create(info *model.Host) error {
	_, err := g.DB().Table(s._host.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Host记录
func (s *host) First(info *request.GetById) (result *model.Host, err error) {
	var entity model.Host
	err = g.DB().Table(s._host.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Host记录
func (s *host) Update(info *model.Host) (result *model.Host, err error) {
    _, err = g.DB().Table(s._host.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Host记录
func (s *host) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._host.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Host记录
func (s *host) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._host.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Host记录列表
func (s *host) GetList(info *request.SearchHost) (list *[]model.Host, total int, err error) {
	var hosts []model.Host
	db := g.DB().Table(s._host.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&hosts)
	return &hosts, total, err
}
