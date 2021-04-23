package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var ClusterNodes = new(clusterNodes)

type clusterNodes struct {
	_clusterNodes model.ClusterNodes
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建ClusterNodes记录
func (s *clusterNodes) Create(info *model.ClusterNodes) error {
	_, err := g.DB().Table(s._clusterNodes.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取ClusterNodes记录
func (s *clusterNodes) First(info *request.GetById) (result *model.ClusterNodes, err error) {
	var entity model.ClusterNodes
	err = g.DB().Table(s._clusterNodes.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新ClusterNodes记录
func (s *clusterNodes) Update(info *model.ClusterNodes) (result *model.ClusterNodes, err error) {
    _, err = g.DB().Table(s._clusterNodes.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除ClusterNodes记录
func (s *clusterNodes) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._clusterNodes.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量ClusterNodes记录
func (s *clusterNodes) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._clusterNodes.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取ClusterNodes记录列表
func (s *clusterNodes) GetList(info *request.SearchClusterNodes) (list *[]model.ClusterNodes, total int, err error) {
	var clusterNodess []model.ClusterNodes
	db := g.DB().Table(s._clusterNodes.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&clusterNodess)
	return &clusterNodess, total, err
}
