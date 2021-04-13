package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	moviesFieldNames          = builderx.RawFieldNames(&Movies{})
	moviesRows                = strings.Join(moviesFieldNames, ",")
	moviesRowsExpectAutoSet   = strings.Join(stringx.Remove(moviesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	moviesRowsWithPlaceHolder = strings.Join(stringx.Remove(moviesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheMoviesIdPrefix    = "cache#movies#id#"
	cacheMoviesTitlePrefix = "cache#movies#title#"
)

type (
	MoviesModel interface {
		Insert(data Movies) (sql.Result, error)
		FindOne(id int64) (*Movies, error)
		FindOneByTitle(title string) (*Movies, error)
		Update(data Movies) error
		Delete(id int64) error
	}

	defaultMoviesModel struct {
		sqlc.CachedConn
		table string
	}

	Movies struct {
		CreateTime  time.Time    `db:"create_time"` // 创建时间
		UpdateTime  time.Time    `db:"update_time"`
		DeletedAt   sql.NullTime `db:"deleted_at"`
		Id          int64        `db:"id"`          // 电影id
		Title       string       `db:"title"`       // 标题
		Description string       `db:"description"` // 描述
		Url         string       `db:"url"`         // 链接
	}
)

func NewMoviesModel(conn sqlx.SqlConn, c cache.CacheConf) MoviesModel {
	return &defaultMoviesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`movies`",
	}
}

func (m *defaultMoviesModel) Insert(data Movies) (sql.Result, error) {
	moviesTitleKey := fmt.Sprintf("%s%v", cacheMoviesTitlePrefix, data.Title)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, moviesRowsExpectAutoSet)
		return conn.Exec(query, data.DeletedAt, data.Title, data.Description, data.Url)
	}, moviesTitleKey)
	return ret, err
}

func (m *defaultMoviesModel) FindOne(id int64) (*Movies, error) {
	moviesIdKey := fmt.Sprintf("%s%v", cacheMoviesIdPrefix, id)
	var resp Movies
	err := m.QueryRow(&resp, moviesIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", moviesRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMoviesModel) FindOneByTitle(title string) (*Movies, error) {
	moviesTitleKey := fmt.Sprintf("%s%v", cacheMoviesTitlePrefix, title)
	var resp Movies
	err := m.QueryRowIndex(&resp, moviesTitleKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `title` = ? limit 1", moviesRows, m.table)
		if err := conn.QueryRow(&resp, query, title); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMoviesModel) Update(data Movies) error {
	moviesIdKey := fmt.Sprintf("%s%v", cacheMoviesIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, moviesRowsWithPlaceHolder)
		return conn.Exec(query, data.DeletedAt, data.Title, data.Description, data.Url, data.Id)
	}, moviesIdKey)
	return err
}

func (m *defaultMoviesModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	moviesIdKey := fmt.Sprintf("%s%v", cacheMoviesIdPrefix, id)
	moviesTitleKey := fmt.Sprintf("%s%v", cacheMoviesTitlePrefix, data.Title)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, moviesIdKey, moviesTitleKey)
	return err
}

func (m *defaultMoviesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMoviesIdPrefix, primary)
}

func (m *defaultMoviesModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", moviesRows, m.table)
	return conn.QueryRow(v, query, primary)
}
