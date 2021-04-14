package svc

import (
	"datacenter/movies/rpc/internal/config"
	"datacenter/movies/model"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)
type ServiceContext struct {
	Config config.Config
	MoviesModel		model.MoviesModel
	RedisConn     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	
	mm := model.NewMoviesModel(conn, c.CacheRedis)
	rconn := redis.NewRedis(c.CacheRedis[0].Host, c.CacheRedis[0].Type, c.CacheRedis[0].Pass)

	return &ServiceContext{
		Config: c,
		MoviesModel: mm,
		RedisConn: rconn,
	}
}
