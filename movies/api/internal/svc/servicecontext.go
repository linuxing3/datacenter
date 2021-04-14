package svc

import (
	"datacenter/movies/api/internal/config"
	"datacenter/movies/api/internal/middleware"
	"datacenter/movies/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	// TODO: 这里可以添加中间件，Rpc服务，拦截器，缓存等等其他服务
	// 参考 internal/svc/servicecontext.go
	// CommonRpc        commonclient.Common 
	// Cache            cache.Cache
	// RedisConn        *redis.Redis
	Config    config.Config
	MoviesModel model.MoviesModel
	Usercheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	mm := model.NewMoviesModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		MoviesModel: mm,
		Usercheck: middleware.NewUsercheckMiddleware().Handle,
	}
}
