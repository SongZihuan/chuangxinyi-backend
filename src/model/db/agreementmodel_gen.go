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
	agreementFieldNames          = builder.RawFieldNames(&Agreement{})
	agreementRows                = strings.Join(agreementFieldNames, ",")
	agreementRowsExpectAutoSet   = strings.Join(stringx.Remove(agreementFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	agreementRowsWithPlaceHolder = strings.Join(stringx.Remove(agreementFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	agreementModel interface {
		Insert(ctx context.Context, data *Agreement) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Agreement, error)
		Update(ctx context.Context, data *Agreement) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAgreementModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Agreement struct {
		Id       int64        `db:"id"`
		Aid      string       `db:"aid"`
		Content  string       `db:"content"`
		CreateAt time.Time    `db:"create_at"`
		UpdateAt time.Time    `db:"update_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
	}
)

func newAgreementModel(conn sqlx.SqlConn) *defaultAgreementModel {
	return &defaultAgreementModel{
		conn:  conn,
		table: "`agreement`",
	}
}

func (m *defaultAgreementModel) withSession(session sqlx.Session) *defaultAgreementModel {
	return &defaultAgreementModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`agreement`",
	}
}

func (m *defaultAgreementModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAgreementModel) FindOne(ctx context.Context, id int64) (*Agreement, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", agreementRows, m.table)
	var resp Agreement
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

func (m *defaultAgreementModel) Insert(ctx context.Context, data *Agreement) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, agreementRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Aid, data.Content, data.DeleteAt)
	return ret, err
}

func (m *defaultAgreementModel) Update(ctx context.Context, data *Agreement) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, agreementRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Aid, data.Content, data.DeleteAt, data.Id)
	return err
}

func (m *defaultAgreementModel) tableName() string {
	return m.table
}
