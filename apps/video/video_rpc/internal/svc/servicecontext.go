package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlConf.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(conn, c.RedisCacheConf),
	}
}
