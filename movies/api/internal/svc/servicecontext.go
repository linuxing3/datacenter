package svc

import (
	"datacenter/movies/api/internal/config"
	"datacenter/movies/api/internal/middleware"
	"github.com/tal-tech/go-zero/rest"
	"datacenter/movies/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	MovieModel *model.DefaultMoviesModel
	Usercheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	mm := model.NewMoviesModel(conn)
	return &ServiceContext{
		Config:    c,
		MovieModel: mm,
		Usercheck: middleware.NewUsercheckMiddleware().Handle,
	}
}
