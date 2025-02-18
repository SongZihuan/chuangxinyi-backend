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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByUid(ctx context.Context, uid string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id              int64         `db:"id"`
		Uid             string        `db:"uid"`
		Status          int64         `db:"status"`
		Signin          bool          `db:"signin"`
		SonLevel        int64         `db:"son_level"`
		FatherId        sql.NullInt64 `db:"father_id"`
		RootFatherId    sql.NullInt64 `db:"root_father_id"`
		InviteId        sql.NullInt64 `db:"invite_id"`
		WalletId        int64         `db:"wallet_id"`
		TokenExpiration int64         `db:"token_expiration"`
		RoleId          int64         `db:"role_id"`
		IsAdmin         bool          `db:"is_admin"`
		Remark          string        `db:"remark"`
		CreateAt        time.Time     `db:"create_at"`
		UpdateAt        time.Time     `db:"update_at"`
		DeleteAt        sql.NullTime  `db:"delete_at"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) withSession(session sqlx.Session) *defaultUserModel {
	return &defaultUserModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
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

func (m *defaultUserModel) FindOneByUid(ctx context.Context, uid string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.Status, data.Signin, data.SonLevel, data.FatherId, data.RootFatherId, data.InviteId, data.WalletId, data.TokenExpiration, data.RoleId, data.IsAdmin, data.Remark, data.DeleteAt)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Uid, newData.Status, newData.Signin, newData.SonLevel, newData.FatherId, newData.RootFatherId, newData.InviteId, newData.WalletId, newData.TokenExpiration, newData.RoleId, newData.IsAdmin, newData.Remark, newData.DeleteAt, newData.Id)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
