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
	withdrawFieldNames          = builder.RawFieldNames(&Withdraw{})
	withdrawRows                = strings.Join(withdrawFieldNames, ",")
	withdrawRowsExpectAutoSet   = strings.Join(stringx.Remove(withdrawFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	withdrawRowsWithPlaceHolder = strings.Join(stringx.Remove(withdrawFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	withdrawModel interface {
		Insert(ctx context.Context, data *Withdraw) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Withdraw, error)
		Update(ctx context.Context, data *Withdraw) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWithdrawModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Withdraw struct {
		Id                int64          `db:"id"`
		WalletId          int64          `db:"wallet_id"`
		UserId            int64          `db:"user_id"`
		WithdrawId        string         `db:"withdraw_id"`
		WithdrawWay       string         `db:"withdraw_way"`
		Name              string         `db:"name"`
		AlipayLoginId     sql.NullString `db:"alipay_login_id"`
		WechatpayOpenId   sql.NullString `db:"wechatpay_open_id"`
		WechatpayUnionId  sql.NullString `db:"wechatpay_union_id"`
		WechatpayNickname sql.NullString `db:"wechatpay_nickname"`
		Cny               int64          `db:"cny"`
		Balance           sql.NullInt64  `db:"balance"`
		OrderId           sql.NullString `db:"order_id"`
		PayFundOrderId    sql.NullString `db:"pay_fund_order_id"`
		Remark            string         `db:"remark"`
		Status            int64          `db:"status"`
		CreateAt          time.Time      `db:"create_at"`
		WithdrawAt        time.Time      `db:"withdraw_at"`
		PayAt             sql.NullTime   `db:"pay_at"`
		DeleteAt          sql.NullTime   `db:"delete_at"`
	}
)

func newWithdrawModel(conn sqlx.SqlConn) *defaultWithdrawModel {
	return &defaultWithdrawModel{
		conn:  conn,
		table: "`withdraw`",
	}
}

func (m *defaultWithdrawModel) withSession(session sqlx.Session) *defaultWithdrawModel {
	return &defaultWithdrawModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`withdraw`",
	}
}

func (m *defaultWithdrawModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultWithdrawModel) FindOne(ctx context.Context, id int64) (*Withdraw, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", withdrawRows, m.table)
	var resp Withdraw
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

func (m *defaultWithdrawModel) Insert(ctx context.Context, data *Withdraw) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, withdrawRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.UserId, data.WithdrawId, data.WithdrawWay, data.Name, data.AlipayLoginId, data.WechatpayOpenId, data.WechatpayUnionId, data.WechatpayNickname, data.Cny, data.Balance, data.OrderId, data.PayFundOrderId, data.Remark, data.Status, data.WithdrawAt, data.PayAt, data.DeleteAt)
	return ret, err
}

func (m *defaultWithdrawModel) Update(ctx context.Context, data *Withdraw) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, withdrawRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.UserId, data.WithdrawId, data.WithdrawWay, data.Name, data.AlipayLoginId, data.WechatpayOpenId, data.WechatpayUnionId, data.WechatpayNickname, data.Cny, data.Balance, data.OrderId, data.PayFundOrderId, data.Remark, data.Status, data.WithdrawAt, data.PayAt, data.DeleteAt, data.Id)
	return err
}

func (m *defaultWithdrawModel) tableName() string {
	return m.table
}