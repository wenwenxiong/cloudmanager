package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var ClusterSpec = new(clusterSpec)

type clusterSpec struct {
	_clusterSpec model.ClusterSpec
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建ClusterSpec记录
func (s *clusterSpec) Create(info *model.ClusterSpec) error {
	_, err := g.DB().Table(s._clusterSpec.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取ClusterSpec记录
func (s *clusterSpec) First(info *request.GetById) (result *model.ClusterSpec, err error) {
	var entity model.ClusterSpec
	err = g.DB().Table(s._clusterSpec.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新ClusterSpec记录
func (s *clusterSpec) Update(info *model.ClusterSpec) (result *model.ClusterSpec, err error) {
    _, err = g.DB().Table(s._clusterSpec.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除ClusterSpec记录
func (s *clusterSpec) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._clusterSpec.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量ClusterSpec记录
func (s *clusterSpec) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._clusterSpec.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取ClusterSpec记录列表
func (s *clusterSpec) GetList(info *request.SearchClusterSpec) (list *[]model.ClusterSpec, total int, err error) {
	var clusterSpecs []model.ClusterSpec
	db := g.DB().Table(s._clusterSpec.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&clusterSpecs)
	return &clusterSpecs, total, err
}
