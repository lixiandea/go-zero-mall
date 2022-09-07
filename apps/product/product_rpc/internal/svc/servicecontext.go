package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlDataSource.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c, c.RedisCacheConf),
	}
}
