package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/order/admin/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/order/order/order"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
