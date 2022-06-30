package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/order/order/order"
	"github.com/lixiandea/go-zero-mall/apps/pay/model"
	"github.com/lixiandea/go-zero-mall/apps/pay/pay/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/user/user/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  user.User
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.RedisCache),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
