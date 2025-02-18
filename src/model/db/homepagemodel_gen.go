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
	homepageFieldNames          = builder.RawFieldNames(&Homepage{})
	homepageRows                = strings.Join(homepageFieldNames, ",")
	homepageRowsExpectAutoSet   = strings.Join(stringx.Remove(homepageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	homepageRowsWithPlaceHolder = strings.Join(stringx.Remove(homepageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	homepageModel interface {
		Insert(ctx context.Context, data *Homepage) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Homepage, error)
		Update(ctx context.Context, data *Homepage) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHomepageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Homepage struct {
		Id           int64          `db:"id"`
		UserId       int64          `db:"user_id"`
		Introduction sql.NullString `db:"introduction"`
		Address      sql.NullString `db:"address"`
		Phone        sql.NullString `db:"phone"`
		Email        sql.NullString `db:"email"`
		Wechat       sql.NullString `db:"wechat"`
		Qq           sql.NullString `db:"qq"`
		Man          sql.NullBool   `db:"man"`
		Link         sql.NullString `db:"link"`
		Company      sql.NullString `db:"company"`
		Industry     sql.NullString `db:"industry"`
		Position     sql.NullString `db:"position"`
		Close        bool           `db:"close"`
		CreateAt     time.Time      `db:"create_at"`
		DeleteAt     sql.NullTime   `db:"delete_at"`
	}
)

func newHomepageModel(conn sqlx.SqlConn) *defaultHomepageModel {
	return &defaultHomepageModel{
		conn:  conn,
		table: "`homepage`",
	}
}

func (m *defaultHomepageModel) withSession(session sqlx.Session) *defaultHomepageModel {
	return &defaultHomepageModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`homepage`",
	}
}

func (m *defaultHomepageModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultHomepageModel) FindOne(ctx context.Context, id int64) (*Homepage, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", homepageRows, m.table)
	var resp Homepage
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

func (m *defaultHomepageModel) Insert(ctx context.Context, data *Homepage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homepageRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Introduction, data.Address, data.Phone, data.Email, data.Wechat, data.Qq, data.Man, data.Link, data.Company, data.Industry, data.Position, data.Close, data.DeleteAt)
	return ret, err
}

func (m *defaultHomepageModel) Update(ctx context.Context, data *Homepage) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homepageRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Introduction, data.Address, data.Phone, data.Email, data.Wechat, data.Qq, data.Man, data.Link, data.Company, data.Industry, data.Position, data.Close, data.DeleteAt, data.Id)
	return err
}

func (m *defaultHomepageModel) tableName() string {
	return m.table
}
