package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRpc      zrpc.RpcClientConf
	CommonRpc    zrpc.RpcClientConf
	VotesRpc     zrpc.RpcClientConf
	SearchRpc    zrpc.RpcClientConf
	QuestionsRpc zrpc.RpcClientConf
	// FIXED: add movie rpc
	MoviesRpc 	 zrpc.RpcClientConf

	CacheRedis cache.ClusterConf
}
