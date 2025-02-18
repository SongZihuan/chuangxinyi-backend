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
	wechatFieldNames          = builder.RawFieldNames(&Wechat{})
	wechatRows                = strings.Join(wechatFieldNames, ",")
	wechatRowsExpectAutoSet   = strings.Join(stringx.Remove(wechatFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	wechatRowsWithPlaceHolder = strings.Join(stringx.Remove(wechatFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	wechatModel interface {
		Insert(ctx context.Context, data *Wechat) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Wechat, error)
		Update(ctx context.Context, data *Wechat) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWechatModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Wechat struct {
		Id         int64          `db:"id"`
		UserId     int64          `db:"user_id"`
		OpenId     sql.NullString `db:"open_id"`
		UnionId    sql.NullString `db:"union_id"`
		Fuwuhao    sql.NullString `db:"fuwuhao"`
		Nickname   sql.NullString `db:"nickname"`
		Headimgurl sql.NullString `db:"headimgurl"`
		IsDelete   bool           `db:"is_delete"`
		CreateAt   time.Time      `db:"create_at"`
		DeleteAt   sql.NullTime   `db:"delete_at"`
	}
)

func newWechatModel(conn sqlx.SqlConn) *defaultWechatModel {
	return &defaultWechatModel{
		conn:  conn,
		table: "`wechat`",
	}
}

func (m *defaultWechatModel) withSession(session sqlx.Session) *defaultWechatModel {
	return &defaultWechatModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`wechat`",
	}
}

func (m *defaultWechatModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultWechatModel) FindOne(ctx context.Context, id int64) (*Wechat, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wechatRows, m.table)
	var resp Wechat
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

func (m *defaultWechatModel) Insert(ctx context.Context, data *Wechat) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, wechatRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.OpenId, data.UnionId, data.Fuwuhao, data.Nickname, data.Headimgurl, data.IsDelete, data.DeleteAt)
	return ret, err
}

func (m *defaultWechatModel) Update(ctx context.Context, data *Wechat) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wechatRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.OpenId, data.UnionId, data.Fuwuhao, data.Nickname, data.Headimgurl, data.IsDelete, data.DeleteAt, data.Id)
	return err
}

func (m *defaultWechatModel) tableName() string {
	return m.table
}
