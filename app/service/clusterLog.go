package service

import (
	"cloudmanager/app/api/request"
	"cloudmanager/app/model"
	"github.com/gogf/gf/frame/g"
)

var ClusterLog = new(clusterLog)

type clusterLog struct {
	_clusterLog model.ClusterLog
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建ClusterLog记录
func (s *clusterLog) Create(info *model.ClusterLog) error {
	_, err := g.DB().Table(s._clusterLog.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取ClusterLog记录
func (s *clusterLog) First(info *request.GetById) (result *model.ClusterLog, err error) {
	var entity model.ClusterLog
	err = g.DB().Table(s._clusterLog.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新ClusterLog记录
func (s *clusterLog) Update(info *model.ClusterLog) (result *model.ClusterLog, err error) {
    _, err = g.DB().Table(s._clusterLog.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除ClusterLog记录
func (s *clusterLog) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._clusterLog.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量ClusterLog记录
func (s *clusterLog) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._clusterLog.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取ClusterLog记录列表
func (s *clusterLog) GetList(info *request.SearchClusterLog) (list *[]model.ClusterLog, total int, err error) {
	var clusterLogs []model.ClusterLog
	db := g.DB().Table(s._clusterLog.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&clusterLogs)
	return &clusterLogs, total, err
}
