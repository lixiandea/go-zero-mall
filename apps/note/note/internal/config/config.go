package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	UserRpcConf zrpc.RpcClientConf
	Mysql       struct {
		DateSource string
	}
	RedisCacheConf cache.CacheConf
}
