package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var ClusterApp = new(clusterApp)

type clusterApp struct {
	_clusterApp model.ClusterApp
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建ClusterApp记录
func (s *clusterApp) Create(info *model.ClusterApp) error {
	_, err := g.DB().Table(s._clusterApp.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取ClusterApp记录
func (s *clusterApp) First(info *request.GetById) (result *model.ClusterApp, err error) {
	var entity model.ClusterApp
	err = g.DB().Table(s._clusterApp.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新ClusterApp记录
func (s *clusterApp) Update(info *model.ClusterApp) (result *model.ClusterApp, err error) {
    _, err = g.DB().Table(s._clusterApp.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除ClusterApp记录
func (s *clusterApp) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._clusterApp.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量ClusterApp记录
func (s *clusterApp) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._clusterApp.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取ClusterApp记录列表
func (s *clusterApp) GetList(info *request.SearchClusterApp) (list *[]model.ClusterApp, total int, err error) {
	var clusterApps []model.ClusterApp
	db := g.DB().Table(s._clusterApp.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&clusterApps)
	return &clusterApps, total, err
}
