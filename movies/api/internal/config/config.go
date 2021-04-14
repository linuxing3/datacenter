package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.ClusterConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	// TODO: 如果需要调用Rpc，可以在这里添加
}

