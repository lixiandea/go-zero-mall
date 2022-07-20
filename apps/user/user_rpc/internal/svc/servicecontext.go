package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlConf.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.RedisCacheConf),
	}
}
