package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisCacheConf cache.CacheConf
	MysqlConf      struct {
		DataSource string
	}
}
