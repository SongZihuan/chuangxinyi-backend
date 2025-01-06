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
	secondfaFieldNames          = builder.RawFieldNames(&Secondfa{})
	secondfaRows                = strings.Join(secondfaFieldNames, ",")
	secondfaRowsExpectAutoSet   = strings.Join(stringx.Remove(secondfaFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	secondfaRowsWithPlaceHolder = strings.Join(stringx.Remove(secondfaFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	secondfaModel interface {
		Insert(ctx context.Context, data *Secondfa) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Secondfa, error)
		Update(ctx context.Context, data *Secondfa) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSecondfaModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Secondfa struct {
		Id       int64          `db:"id"`
		UserId   int64          `db:"user_id"`
		Secret   sql.NullString `db:"secret"`
		CreateAt time.Time      `db:"create_at"`
		DeleteAt sql.NullTime   `db:"delete_at"`
	}
)

func newSecondfaModel(conn sqlx.SqlConn) *defaultSecondfaModel {
	return &defaultSecondfaModel{
		conn:  conn,
		table: "`secondfa`",
	}
}

func (m *defaultSecondfaModel) withSession(session sqlx.Session) *defaultSecondfaModel {
	return &defaultSecondfaModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`secondfa`",
	}
}

func (m *defaultSecondfaModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSecondfaModel) FindOne(ctx context.Context, id int64) (*Secondfa, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", secondfaRows, m.table)
	var resp Secondfa
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

func (m *defaultSecondfaModel) Insert(ctx context.Context, data *Secondfa) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, secondfaRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Secret, data.DeleteAt)
	return ret, err
}

func (m *defaultSecondfaModel) Update(ctx context.Context, data *Secondfa) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, secondfaRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Secret, data.DeleteAt, data.Id)
	return err
}

func (m *defaultSecondfaModel) tableName() string {
	return m.table
}
