package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/order/model"
	"github.com/lixiandea/go-zero-mall/apps/order/order/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/product/product/product"
	"github.com/lixiandea/go-zero-mall/apps/user/user/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
	UserRpc    user.User
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.RedisCache),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
