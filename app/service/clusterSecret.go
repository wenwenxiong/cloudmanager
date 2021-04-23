package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var ClusterSecret = new(clusterSecret)

type clusterSecret struct {
	_clusterSecret model.ClusterSecret
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建ClusterSecret记录
func (s *clusterSecret) Create(info *model.ClusterSecret) error {
	_, err := g.DB().Table(s._clusterSecret.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取ClusterSecret记录
func (s *clusterSecret) First(info *request.GetById) (result *model.ClusterSecret, err error) {
	var entity model.ClusterSecret
	err = g.DB().Table(s._clusterSecret.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新ClusterSecret记录
func (s *clusterSecret) Update(info *model.ClusterSecret) (result *model.ClusterSecret, err error) {
    _, err = g.DB().Table(s._clusterSecret.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除ClusterSecret记录
func (s *clusterSecret) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._clusterSecret.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量ClusterSecret记录
func (s *clusterSecret) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._clusterSecret.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取ClusterSecret记录列表
func (s *clusterSecret) GetList(info *request.SearchClusterSecret) (list *[]model.ClusterSecret, total int, err error) {
	var clusterSecrets []model.ClusterSecret
	db := g.DB().Table(s._clusterSecret.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&clusterSecrets)
	return &clusterSecrets, total, err
}
