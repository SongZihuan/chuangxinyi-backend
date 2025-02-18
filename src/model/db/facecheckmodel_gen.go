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
	faceCheckFieldNames          = builder.RawFieldNames(&FaceCheck{})
	faceCheckRows                = strings.Join(faceCheckFieldNames, ",")
	faceCheckRowsExpectAutoSet   = strings.Join(stringx.Remove(faceCheckFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	faceCheckRowsWithPlaceHolder = strings.Join(stringx.Remove(faceCheckFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	faceCheckModel interface {
		Insert(ctx context.Context, data *FaceCheck) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FaceCheck, error)
		Update(ctx context.Context, data *FaceCheck) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFaceCheckModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FaceCheck struct {
		Id        int64        `db:"id"`
		CheckId   string       `db:"check_id"`
		CertifyId string       `db:"certify_id"`
		Name      string       `db:"name"`
		Idcard    string       `db:"idcard"`
		Status    int64        `db:"status"`
		CreateAt  time.Time    `db:"create_at"`
		DeleteAt  sql.NullTime `db:"delete_at"`
	}
)

func newFaceCheckModel(conn sqlx.SqlConn) *defaultFaceCheckModel {
	return &defaultFaceCheckModel{
		conn:  conn,
		table: "`face_check`",
	}
}

func (m *defaultFaceCheckModel) withSession(session sqlx.Session) *defaultFaceCheckModel {
	return &defaultFaceCheckModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`face_check`",
	}
}

func (m *defaultFaceCheckModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFaceCheckModel) FindOne(ctx context.Context, id int64) (*FaceCheck, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", faceCheckRows, m.table)
	var resp FaceCheck
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

func (m *defaultFaceCheckModel) Insert(ctx context.Context, data *FaceCheck) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, faceCheckRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CheckId, data.CertifyId, data.Name, data.Idcard, data.Status, data.DeleteAt)
	return ret, err
}

func (m *defaultFaceCheckModel) Update(ctx context.Context, data *FaceCheck) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, faceCheckRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CheckId, data.CertifyId, data.Name, data.Idcard, data.Status, data.DeleteAt, data.Id)
	return err
}

func (m *defaultFaceCheckModel) tableName() string {
	return m.table
}
