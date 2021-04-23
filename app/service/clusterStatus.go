package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var ClusterStatus = new(clusterStatus)

type clusterStatus struct {
	_clusterStatus model.ClusterStatus
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建ClusterStatus记录
func (s *clusterStatus) Create(info *model.ClusterStatus) error {
	_, err := g.DB().Table(s._clusterStatus.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取ClusterStatus记录
func (s *clusterStatus) First(info *request.GetById) (result *model.ClusterStatus, err error) {
	var entity model.ClusterStatus
	err = g.DB().Table(s._clusterStatus.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新ClusterStatus记录
func (s *clusterStatus) Update(info *model.ClusterStatus) (result *model.ClusterStatus, err error) {
    _, err = g.DB().Table(s._clusterStatus.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除ClusterStatus记录
func (s *clusterStatus) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._clusterStatus.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量ClusterStatus记录
func (s *clusterStatus) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._clusterStatus.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取ClusterStatus记录列表
func (s *clusterStatus) GetList(info *request.SearchClusterStatus) (list *[]model.ClusterStatus, total int, err error) {
	var clusterStatuss []model.ClusterStatus
	db := g.DB().Table(s._clusterStatus.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&clusterStatuss)
	return &clusterStatuss, total, err
}
