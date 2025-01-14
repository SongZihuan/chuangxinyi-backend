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
	auditFieldNames          = builder.RawFieldNames(&Audit{})
	auditRows                = strings.Join(auditFieldNames, ",")
	auditRowsExpectAutoSet   = strings.Join(stringx.Remove(auditFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	auditRowsWithPlaceHolder = strings.Join(stringx.Remove(auditFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	auditModel interface {
		Insert(ctx context.Context, data *Audit) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Audit, error)
		Update(ctx context.Context, data *Audit) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAuditModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Audit struct {
		Id       int64        `db:"id"`
		UserId   int64        `db:"user_id"`
		Content  string       `db:"content"`
		From     string       `db:"from"`
		FromId   int64        `db:"from_id"`
		CreateAt time.Time    `db:"create_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
	}
)

func newAuditModel(conn sqlx.SqlConn) *defaultAuditModel {
	return &defaultAuditModel{
		conn:  conn,
		table: "`audit`",
	}
}

func (m *defaultAuditModel) withSession(session sqlx.Session) *defaultAuditModel {
	return &defaultAuditModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`audit`",
	}
}

func (m *defaultAuditModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAuditModel) FindOne(ctx context.Context, id int64) (*Audit, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", auditRows, m.table)
	var resp Audit
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

func (m *defaultAuditModel) Insert(ctx context.Context, data *Audit) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, auditRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Content, data.From, data.FromId, data.DeleteAt)
	return ret, err
}

func (m *defaultAuditModel) Update(ctx context.Context, data *Audit) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, auditRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Content, data.From, data.FromId, data.DeleteAt, data.Id)
	return err
}

func (m *defaultAuditModel) tableName() string {
	return m.table
}
