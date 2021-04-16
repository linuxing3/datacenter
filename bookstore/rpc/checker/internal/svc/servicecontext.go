package svc

import (
	"datacenter/bookstore/model"
	"datacenter/bookstore/rpc/checker/internal/config"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	BookModel model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BookModel: model.NewBookModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
