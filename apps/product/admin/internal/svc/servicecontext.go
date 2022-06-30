package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/product/admin/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/product/rpc/rpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc rpc.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: rpc.NewRpc(zrpc.MustNewClient(c.ProductRpc)),
	}
}
