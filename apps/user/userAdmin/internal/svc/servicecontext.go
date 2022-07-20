package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/userrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
