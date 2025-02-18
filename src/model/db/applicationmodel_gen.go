// Code generated by goctlwt. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/wuntsong-org/go-zero-plus/core/stores/builder"
	"github.com/wuntsong-org/go-zero-plus/core/stores/sqlc"
	"github.com/wuntsong-org/go-zero-plus/core/stores/sqlx"
	"github.com/wuntsong-org/go-zero-plus/core/stringx"
)

var (
	applicationFieldNames          = builder.RawFieldNames(&Application{})
	applicationRows                = strings.Join(applicationFieldNames, ",")
	applicationRowsExpectAutoSet   = strings.Join(stringx.Remove(applicationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	applicationRowsWithPlaceHolder = strings.Join(stringx.Remove(applicationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	applicationModel interface {
		Insert(ctx context.Context, data *Application) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Application, error)
		Update(ctx context.Context, data *Application) error
		Delete(ctx context.Context, id int64) error
	}

	defaultApplicationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Application struct {
		Id       int64        `db:"id"`
		Name     string       `db:"name"`
		Describe string       `db:"describe"`
		WebId    int64        `db:"web_id"`
		Url      string       `db:"url"`
		Icon     string       `db:"icon"`
		Status   int64        `db:"status"`
		Sort     int64        `db:"sort"`
		CreateAt time.Time    `db:"create_at"`
		UpdateAt time.Time    `db:"update_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
	}
)

func newApplicationModel(conn sqlx.SqlConn) *defaultApplicationModel {
	return &defaultApplicationModel{
		conn:  conn,
		table: "`application`",
	}
}

func (m *defaultApplicationModel) withSession(session sqlx.Session) *defaultApplicationModel {
	return &defaultApplicationModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`application`",
	}
}

func (m *defaultApplicationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultApplicationModel) FindOne(ctx context.Context, id int64) (*Application, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationRows, m.table)
	var resp Application
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultApplicationModel) Insert(ctx context.Context, data *Application) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, applicationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Describe, data.WebId, data.Url, data.Icon, data.Status, data.Sort, data.DeleteAt)
	return ret, err
}

func (m *defaultApplicationModel) Update(ctx context.Context, data *Application) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, applicationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Describe, data.WebId, data.Url, data.Icon, data.Status, data.Sort, data.DeleteAt, data.Id)
	return err
}

func (m *defaultApplicationModel) tableName() string {
	return m.table
}
