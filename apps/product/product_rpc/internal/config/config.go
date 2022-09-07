package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlDataSource struct {
		DataSource string
	}
	RedisCacheConf cache.CacheConf
}
