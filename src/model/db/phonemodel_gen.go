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
	phoneFieldNames          = builder.RawFieldNames(&Phone{})
	phoneRows                = strings.Join(phoneFieldNames, ",")
	phoneRowsExpectAutoSet   = strings.Join(stringx.Remove(phoneFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	phoneRowsWithPlaceHolder = strings.Join(stringx.Remove(phoneFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	phoneModel interface {
		Insert(ctx context.Context, data *Phone) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Phone, error)
		Update(ctx context.Context, data *Phone) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPhoneModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Phone struct {
		Id       int64        `db:"id"`
		UserId   int64        `db:"user_id"`
		Phone    string       `db:"phone"`
		IsDelete bool         `db:"is_delete"`
		CreateAt time.Time    `db:"create_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
	}
)

func newPhoneModel(conn sqlx.SqlConn) *defaultPhoneModel {
	return &defaultPhoneModel{
		conn:  conn,
		table: "`phone`",
	}
}

func (m *defaultPhoneModel) withSession(session sqlx.Session) *defaultPhoneModel {
	return &defaultPhoneModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`phone`",
	}
}

func (m *defaultPhoneModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPhoneModel) FindOne(ctx context.Context, id int64) (*Phone, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", phoneRows, m.table)
	var resp Phone
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

func (m *defaultPhoneModel) Insert(ctx context.Context, data *Phone) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, phoneRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Phone, data.IsDelete, data.DeleteAt)
	return ret, err
}

func (m *defaultPhoneModel) Update(ctx context.Context, data *Phone) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, phoneRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Phone, data.IsDelete, data.DeleteAt, data.Id)
	return err
}

func (m *defaultPhoneModel) tableName() string {
	return m.table
}
