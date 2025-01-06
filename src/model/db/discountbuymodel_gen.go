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
	discountBuyFieldNames          = builder.RawFieldNames(&DiscountBuy{})
	discountBuyRows                = strings.Join(discountBuyFieldNames, ",")
	discountBuyRowsExpectAutoSet   = strings.Join(stringx.Remove(discountBuyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	discountBuyRowsWithPlaceHolder = strings.Join(stringx.Remove(discountBuyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	discountBuyModel interface {
		Insert(ctx context.Context, data *DiscountBuy) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DiscountBuy, error)
		Update(ctx context.Context, data *DiscountBuy) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDiscountBuyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DiscountBuy struct {
		Id            int64        `db:"id"`
		UserId        int64        `db:"user_id"`
		DiscountId    int64        `db:"discount_id"`
		Name          string       `db:"name"`
		ShortDescribe string       `db:"short_describe"`
		Days          int64        `db:"days"`
		Month         int64        `db:"month"`
		Year          int64        `db:"year"`
		All           int64        `db:"all"`
		CreateAt      time.Time    `db:"create_at"`
		DeleteAt      sql.NullTime `db:"delete_at"`
	}
)

func newDiscountBuyModel(conn sqlx.SqlConn) *defaultDiscountBuyModel {
	return &defaultDiscountBuyModel{
		conn:  conn,
		table: "`discount_buy`",
	}
}

func (m *defaultDiscountBuyModel) withSession(session sqlx.Session) *defaultDiscountBuyModel {
	return &defaultDiscountBuyModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`discount_buy`",
	}
}

func (m *defaultDiscountBuyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDiscountBuyModel) FindOne(ctx context.Context, id int64) (*DiscountBuy, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", discountBuyRows, m.table)
	var resp DiscountBuy
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

func (m *defaultDiscountBuyModel) Insert(ctx context.Context, data *DiscountBuy) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, discountBuyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.DiscountId, data.Name, data.ShortDescribe, data.Days, data.Month, data.Year, data.All, data.DeleteAt)
	return ret, err
}

func (m *defaultDiscountBuyModel) Update(ctx context.Context, data *DiscountBuy) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, discountBuyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.DiscountId, data.Name, data.ShortDescribe, data.Days, data.Month, data.Year, data.All, data.DeleteAt, data.Id)
	return err
}

func (m *defaultDiscountBuyModel) tableName() string {
	return m.table
}