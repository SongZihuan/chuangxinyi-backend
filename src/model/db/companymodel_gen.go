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
	companyFieldNames          = builder.RawFieldNames(&Company{})
	companyRows                = strings.Join(companyFieldNames, ",")
	companyRowsExpectAutoSet   = strings.Join(stringx.Remove(companyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	companyRowsWithPlaceHolder = strings.Join(stringx.Remove(companyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	companyModel interface {
		Insert(ctx context.Context, data *Company) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Company, error)
		Update(ctx context.Context, data *Company) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCompanyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Company struct {
		Id                int64          `db:"id"`
		UserId            int64          `db:"user_id"`
		LegalPersonName   string         `db:"legal_person_name"`
		LegalPersonIdCard string         `db:"legal_person_id_card"`
		CompanyName       string         `db:"company_name"`
		CompanyId         string         `db:"company_id"`
		LicenseKey        sql.NullString `db:"license_key"`
		IdcardKey         sql.NullString `db:"idcard_key"`
		IdcardBackKey     sql.NullString `db:"idcard_back_key"`
		FaceCheckId       sql.NullString `db:"face_check_id"`
		IsDelete          bool           `db:"is_delete"`
		CreateAt          time.Time      `db:"create_at"`
		UpdateAt          time.Time      `db:"update_at"`
		DeleteAt          sql.NullTime   `db:"delete_at"`
	}
)

func newCompanyModel(conn sqlx.SqlConn) *defaultCompanyModel {
	return &defaultCompanyModel{
		conn:  conn,
		table: "`company`",
	}
}

func (m *defaultCompanyModel) withSession(session sqlx.Session) *defaultCompanyModel {
	return &defaultCompanyModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`company`",
	}
}

func (m *defaultCompanyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCompanyModel) FindOne(ctx context.Context, id int64) (*Company, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", companyRows, m.table)
	var resp Company
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

func (m *defaultCompanyModel) Insert(ctx context.Context, data *Company) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, companyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.LegalPersonName, data.LegalPersonIdCard, data.CompanyName, data.CompanyId, data.LicenseKey, data.IdcardKey, data.IdcardBackKey, data.FaceCheckId, data.IsDelete, data.DeleteAt)
	return ret, err
}

func (m *defaultCompanyModel) Update(ctx context.Context, data *Company) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, companyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.LegalPersonName, data.LegalPersonIdCard, data.CompanyName, data.CompanyId, data.LicenseKey, data.IdcardKey, data.IdcardBackKey, data.FaceCheckId, data.IsDelete, data.DeleteAt, data.Id)
	return err
}

func (m *defaultCompanyModel) tableName() string {
	return m.table
}