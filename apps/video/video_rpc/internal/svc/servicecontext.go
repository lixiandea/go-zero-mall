package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/userrpc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	UserRpc    user_rpc.UserRpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlConf.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(conn, c.RedisCacheConf),
		UserRpc:    userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
