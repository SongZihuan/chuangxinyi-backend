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
	websitePolicyFieldNames          = builder.RawFieldNames(&WebsitePolicy{})
	websitePolicyRows                = strings.Join(websitePolicyFieldNames, ",")
	websitePolicyRowsExpectAutoSet   = strings.Join(stringx.Remove(websitePolicyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	websitePolicyRowsWithPlaceHolder = strings.Join(stringx.Remove(websitePolicyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	websitePolicyModel interface {
		Insert(ctx context.Context, data *WebsitePolicy) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*WebsitePolicy, error)
		FindOneById(ctx context.Context, id int64) (*WebsitePolicy, error)
		Update(ctx context.Context, data *WebsitePolicy) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWebsitePolicyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	WebsitePolicy struct {
		Id       int64        `db:"id"`
		Name     string       `db:"name"`
		Sign     string       `db:"sign"`
		Describe string       `db:"describe"`
		Sort     int64        `db:"sort"`
		Status   int64        `db:"status"`
		CreateAt time.Time    `db:"create_at"`
		UpdateAt time.Time    `db:"update_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
	}
)

func newWebsitePolicyModel(conn sqlx.SqlConn) *defaultWebsitePolicyModel {
	return &defaultWebsitePolicyModel{
		conn:  conn,
		table: "`website_policy`",
	}
}

func (m *defaultWebsitePolicyModel) withSession(session sqlx.Session) *defaultWebsitePolicyModel {
	return &defaultWebsitePolicyModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`website_policy`",
	}
}

func (m *defaultWebsitePolicyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultWebsitePolicyModel) FindOne(ctx context.Context, id int64) (*WebsitePolicy, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", websitePolicyRows, m.table)
	var resp WebsitePolicy
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

func (m *defaultWebsitePolicyModel) FindOneById(ctx context.Context, id int64) (*WebsitePolicy, error) {
	var resp WebsitePolicy
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", websitePolicyRows, m.table)
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

func (m *defaultWebsitePolicyModel) Insert(ctx context.Context, data *WebsitePolicy) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, websitePolicyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Sign, data.Describe, data.Sort, data.Status, data.DeleteAt)
	return ret, err
}

func (m *defaultWebsitePolicyModel) Update(ctx context.Context, newData *WebsitePolicy) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, websitePolicyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.Sign, newData.Describe, newData.Sort, newData.Status, newData.DeleteAt, newData.Id)
	return err
}

func (m *defaultWebsitePolicyModel) tableName() string {
	return m.table
}