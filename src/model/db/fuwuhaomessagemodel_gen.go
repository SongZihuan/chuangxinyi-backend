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
	fuwuhaoMessageFieldNames          = builder.RawFieldNames(&FuwuhaoMessage{})
	fuwuhaoMessageRows                = strings.Join(fuwuhaoMessageFieldNames, ",")
	fuwuhaoMessageRowsExpectAutoSet   = strings.Join(stringx.Remove(fuwuhaoMessageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	fuwuhaoMessageRowsWithPlaceHolder = strings.Join(stringx.Remove(fuwuhaoMessageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	fuwuhaoMessageModel interface {
		Insert(ctx context.Context, data *FuwuhaoMessage) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FuwuhaoMessage, error)
		Update(ctx context.Context, data *FuwuhaoMessage) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFuwuhaoMessageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FuwuhaoMessage struct {
		Id       int64          `db:"id"`
		OpenId   string         `db:"open_id"`
		Template string         `db:"template"`
		Url      string         `db:"url"`
		Val      string         `db:"val"`
		SenderId int64          `db:"sender_id"`
		Success  bool           `db:"success"`
		ErrorMsg sql.NullString `db:"error_msg"`
		CreateAt time.Time      `db:"create_at"`
		DeleteAt sql.NullTime   `db:"delete_at"`
	}
)

func newFuwuhaoMessageModel(conn sqlx.SqlConn) *defaultFuwuhaoMessageModel {
	return &defaultFuwuhaoMessageModel{
		conn:  conn,
		table: "`fuwuhao_message`",
	}
}

func (m *defaultFuwuhaoMessageModel) withSession(session sqlx.Session) *defaultFuwuhaoMessageModel {
	return &defaultFuwuhaoMessageModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`fuwuhao_message`",
	}
}

func (m *defaultFuwuhaoMessageModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFuwuhaoMessageModel) FindOne(ctx context.Context, id int64) (*FuwuhaoMessage, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", fuwuhaoMessageRows, m.table)
	var resp FuwuhaoMessage
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

func (m *defaultFuwuhaoMessageModel) Insert(ctx context.Context, data *FuwuhaoMessage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, fuwuhaoMessageRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OpenId, data.Template, data.Url, data.Val, data.SenderId, data.Success, data.ErrorMsg, data.DeleteAt)
	return ret, err
}

func (m *defaultFuwuhaoMessageModel) Update(ctx context.Context, data *FuwuhaoMessage) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, fuwuhaoMessageRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OpenId, data.Template, data.Url, data.Val, data.SenderId, data.Success, data.ErrorMsg, data.DeleteAt, data.Id)
	return err
}

func (m *defaultFuwuhaoMessageModel) tableName() string {
	return m.table
}
