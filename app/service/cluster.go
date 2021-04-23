package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var Cluster = new(cluster)

type cluster struct {
	_cluster model.Cluster
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建Cluster记录
func (s *cluster) Create(info *model.Cluster) error {
	_, err := g.DB().Table(s._cluster.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取Cluster记录
func (s *cluster) First(info *request.GetById) (result *model.Cluster, err error) {
	var entity model.Cluster
	err = g.DB().Table(s._cluster.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新Cluster记录
func (s *cluster) Update(info *model.Cluster) (result *model.Cluster, err error) {
    _, err = g.DB().Table(s._cluster.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除Cluster记录
func (s *cluster) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._cluster.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量Cluster记录
func (s *cluster) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._cluster.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取Cluster记录列表
func (s *cluster) GetList(info *request.SearchCluster) (list *[]model.Cluster, total int, err error) {
	var clusters []model.Cluster
	db := g.DB().Table(s._cluster.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&clusters)
	return &clusters, total, err
}
