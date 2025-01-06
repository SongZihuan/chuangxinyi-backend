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
	payFieldNames          = builder.RawFieldNames(&Pay{})
	payRows                = strings.Join(payFieldNames, ",")
	payRowsExpectAutoSet   = strings.Join(stringx.Remove(payFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	payRowsWithPlaceHolder = strings.Join(stringx.Remove(payFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	payModel interface {
		Insert(ctx context.Context, data *Pay) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Pay, error)
		Update(ctx context.Context, data *Pay) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPayModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Pay struct {
		Id          int64          `db:"id"`
		WalletId    int64          `db:"wallet_id"`
		UserId      int64          `db:"user_id"`
		Subject     string         `db:"subject"`
		PayWay      string         `db:"pay_way"`
		PayId       string         `db:"pay_id"`
		Cny         int64          `db:"cny"`
		Get         int64          `db:"get"`
		CouponsId   sql.NullInt64  `db:"coupons_id"`
		TradeNo     sql.NullString `db:"trade_no"`
		BuyerId     sql.NullString `db:"buyer_id"`
		TradeStatus int64          `db:"trade_status"`
		Balance     sql.NullInt64  `db:"balance"`
		Remark      string         `db:"remark"`
		CreateAt    time.Time      `db:"create_at"`
		PayAt       sql.NullTime   `db:"pay_at"`
		RefundAt    sql.NullTime   `db:"refund_at"`
		DeleteAt    sql.NullTime   `db:"delete_at"`
	}
)

func newPayModel(conn sqlx.SqlConn) *defaultPayModel {
	return &defaultPayModel{
		conn:  conn,
		table: "`pay`",
	}
}

func (m *defaultPayModel) withSession(session sqlx.Session) *defaultPayModel {
	return &defaultPayModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`pay`",
	}
}

func (m *defaultPayModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPayModel) FindOne(ctx context.Context, id int64) (*Pay, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", payRows, m.table)
	var resp Pay
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

func (m *defaultPayModel) Insert(ctx context.Context, data *Pay) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, payRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.UserId, data.Subject, data.PayWay, data.PayId, data.Cny, data.Get, data.CouponsId, data.TradeNo, data.BuyerId, data.TradeStatus, data.Balance, data.Remark, data.PayAt, data.RefundAt, data.DeleteAt)
	return ret, err
}

func (m *defaultPayModel) Update(ctx context.Context, data *Pay) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, payRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.UserId, data.Subject, data.PayWay, data.PayId, data.Cny, data.Get, data.CouponsId, data.TradeNo, data.BuyerId, data.TradeStatus, data.Balance, data.Remark, data.PayAt, data.RefundAt, data.DeleteAt, data.Id)
	return err
}

func (m *defaultPayModel) tableName() string {
	return m.table
}
