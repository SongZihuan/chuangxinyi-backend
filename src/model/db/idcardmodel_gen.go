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
	idcardFieldNames          = builder.RawFieldNames(&Idcard{})
	idcardRows                = strings.Join(idcardFieldNames, ",")
	idcardRowsExpectAutoSet   = strings.Join(stringx.Remove(idcardFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	idcardRowsWithPlaceHolder = strings.Join(stringx.Remove(idcardFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	idcardModel interface {
		Insert(ctx context.Context, data *Idcard) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Idcard, error)
		Update(ctx context.Context, data *Idcard) error
		Delete(ctx context.Context, id int64) error
	}

	defaultIdcardModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Idcard struct {
		Id            int64          `db:"id"`
		UserId        int64          `db:"user_id"`
		UserName      string         `db:"user_name"`
		UserIdCard    string         `db:"user_id_card"`
		Phone         sql.NullString `db:"phone"`
		IdcardKey     sql.NullString `db:"idcard_key"`
		IdcardBackKey sql.NullString `db:"idcard_back_key"`
		FaceCheckId   sql.NullString `db:"face_check_id"`
		IsCompany     bool           `db:"is_company"`
		IsDelete      bool           `db:"is_delete"`
		CreateAt      time.Time      `db:"create_at"`
		UpdateAt      time.Time      `db:"update_at"`
		DeleteAt      sql.NullTime   `db:"delete_at"`
	}
)

func newIdcardModel(conn sqlx.SqlConn) *defaultIdcardModel {
	return &defaultIdcardModel{
		conn:  conn,
		table: "`idcard`",
	}
}

func (m *defaultIdcardModel) withSession(session sqlx.Session) *defaultIdcardModel {
	return &defaultIdcardModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`idcard`",
	}
}

func (m *defaultIdcardModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultIdcardModel) FindOne(ctx context.Context, id int64) (*Idcard, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", idcardRows, m.table)
	var resp Idcard
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

func (m *defaultIdcardModel) Insert(ctx context.Context, data *Idcard) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, idcardRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserName, data.UserIdCard, data.Phone, data.IdcardKey, data.IdcardBackKey, data.FaceCheckId, data.IsCompany, data.IsDelete, data.DeleteAt)
	return ret, err
}

func (m *defaultIdcardModel) Update(ctx context.Context, data *Idcard) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, idcardRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserName, data.UserIdCard, data.Phone, data.IdcardKey, data.IdcardBackKey, data.FaceCheckId, data.IsCompany, data.IsDelete, data.DeleteAt, data.Id)
	return err
}

func (m *defaultIdcardModel) tableName() string {
	return m.table
}
