package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

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
)

type (
	MoviesModel interface {
		Insert(data Movies) (sql.Result, error)
		FindOne(id int64) (*Movies, error)
		FindOneByTitle(title string) (*Movies, error)
		Update(data Movies) error
		Delete(id int64) error
	}

	DefaultMoviesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Movies struct {
		Id          string        `db:"id"`          // 电影id
		Title       string       `db:"title"`       // 标题
		Description string       `db:"description"` // 描述
		Url         string       `db:"url"`         // 链接
		CreateTime  time.Time    `db:"create_time"` // 创建时间
		UpdateTime  time.Time    `db:"update_time"`
		DeletedAt   sql.NullTime `db:"deleted_at"`
	}
)

func NewMoviesModel(conn sqlx.SqlConn) *DefaultMoviesModel {
	return &DefaultMoviesModel{
		conn:  conn,
		table: "`movies`",
	}
}

func (m *DefaultMoviesModel) Insert(data Movies) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, moviesRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Title, data.Description, data.Url, data.DeletedAt)
	return ret, err
}

func (m *DefaultMoviesModel) FindOne(id int64) (*Movies, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", moviesRows, m.table)
	var resp Movies
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *DefaultMoviesModel) FindOneByTitle(title string) (*Movies, error) {
	var resp Movies
	query := fmt.Sprintf("select %s from %s where `title` = ? limit 1", moviesRows, m.table)
	err := m.conn.QueryRow(&resp, query, title)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *DefaultMoviesModel) Update(data Movies) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, moviesRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Title, data.Description, data.Url, data.DeletedAt, data.Id)
	return err
}

func (m *DefaultMoviesModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
