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
	addressFieldNames          = builder.RawFieldNames(&Address{})
	addressRows                = strings.Join(addressFieldNames, ",")
	addressRowsExpectAutoSet   = strings.Join(stringx.Remove(addressFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	addressRowsWithPlaceHolder = strings.Join(stringx.Remove(addressFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	addressModel interface {
		Insert(ctx context.Context, data *Address) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Address, error)
		Update(ctx context.Context, data *Address) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAddressModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Address struct {
		Id           int64          `db:"id"`
		UserId       int64          `db:"user_id"`
		Name         sql.NullString `db:"name"`
		Phone        sql.NullString `db:"phone"`
		Email        sql.NullString `db:"email"`
		Country      sql.NullString `db:"country"`
		Province     sql.NullString `db:"province"`
		City         sql.NullString `db:"city"`
		District     sql.NullString `db:"district"`
		CountryCode  sql.NullString `db:"country_code"`
		ProvinceCode sql.NullString `db:"province_code"`
		CityCode     sql.NullString `db:"city_code"`
		DistrictCode sql.NullString `db:"district_code"`
		Address      sql.NullString `db:"address"`
		CreateAt     time.Time      `db:"create_at"`
		DeleteAt     sql.NullTime   `db:"delete_at"`
	}
)

func newAddressModel(conn sqlx.SqlConn) *defaultAddressModel {
	return &defaultAddressModel{
		conn:  conn,
		table: "`address`",
	}
}

func (m *defaultAddressModel) withSession(session sqlx.Session) *defaultAddressModel {
	return &defaultAddressModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`address`",
	}
}

func (m *defaultAddressModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAddressModel) FindOne(ctx context.Context, id int64) (*Address, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", addressRows, m.table)
	var resp Address
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

func (m *defaultAddressModel) Insert(ctx context.Context, data *Address) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, addressRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Name, data.Phone, data.Email, data.Country, data.Province, data.City, data.District, data.CountryCode, data.ProvinceCode, data.CityCode, data.DistrictCode, data.Address, data.DeleteAt)
	return ret, err
}

func (m *defaultAddressModel) Update(ctx context.Context, data *Address) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, addressRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Name, data.Phone, data.Email, data.Country, data.Province, data.City, data.District, data.CountryCode, data.ProvinceCode, data.CityCode, data.DistrictCode, data.Address, data.DeleteAt, data.Id)
	return err
}

func (m *defaultAddressModel) tableName() string {
	return m.table
}
