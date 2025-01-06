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
	tokenRecordFieldNames          = builder.RawFieldNames(&TokenRecord{})
	tokenRecordRows                = strings.Join(tokenRecordFieldNames, ",")
	tokenRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(tokenRecordFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tokenRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(tokenRecordFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tokenRecordModel interface {
		Insert(ctx context.Context, data *TokenRecord) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TokenRecord, error)
		Update(ctx context.Context, data *TokenRecord) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTokenRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TokenRecord struct {
		Id        int64        `db:"id"`
		TokenType int64        `db:"token_type"`
		Token     string       `db:"token"`
		Type      int64        `db:"type"`
		Data      string       `db:"data"`
		CreateAt  time.Time    `db:"create_at"`
		DeleteAt  sql.NullTime `db:"delete_at"`
	}
)

func newTokenRecordModel(conn sqlx.SqlConn) *defaultTokenRecordModel {
	return &defaultTokenRecordModel{
		conn:  conn,
		table: "`token_record`",
	}
}

func (m *defaultTokenRecordModel) withSession(session sqlx.Session) *defaultTokenRecordModel {
	return &defaultTokenRecordModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`token_record`",
	}
}

func (m *defaultTokenRecordModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTokenRecordModel) FindOne(ctx context.Context, id int64) (*TokenRecord, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tokenRecordRows, m.table)
	var resp TokenRecord
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

func (m *defaultTokenRecordModel) Insert(ctx context.Context, data *TokenRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, tokenRecordRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.TokenType, data.Token, data.Type, data.Data, data.DeleteAt)
	return ret, err
}

func (m *defaultTokenRecordModel) Update(ctx context.Context, data *TokenRecord) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tokenRecordRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.TokenType, data.Token, data.Type, data.Data, data.DeleteAt, data.Id)
	return err
}

func (m *defaultTokenRecordModel) tableName() string {
	return m.table
}
