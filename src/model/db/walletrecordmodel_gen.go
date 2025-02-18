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
	walletRecordFieldNames          = builder.RawFieldNames(&WalletRecord{})
	walletRecordRows                = strings.Join(walletRecordFieldNames, ",")
	walletRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(walletRecordFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	walletRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(walletRecordFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	walletRecordModel interface {
		Insert(ctx context.Context, data *WalletRecord) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*WalletRecord, error)
		Update(ctx context.Context, data *WalletRecord) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWalletRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	WalletRecord struct {
		Id                 int64        `db:"id"`
		WalletId           int64        `db:"wallet_id"`
		UserId             int64        `db:"user_id"`
		Type               int64        `db:"type"`
		FundingId          string       `db:"funding_id"`
		Reason             string       `db:"reason"`
		Balance            int64        `db:"balance"`
		WaitBalance        int64        `db:"wait_balance"`
		Cny                int64        `db:"cny"`
		NotBilled          int64        `db:"not_billed"`
		Billed             int64        `db:"billed"`
		HasBilled          int64        `db:"has_billed"`
		Withdraw           int64        `db:"withdraw"`
		WaitWithdraw       int64        `db:"wait_withdraw"`
		NotWithdraw        int64        `db:"not_withdraw"`
		HasWithdraw        int64        `db:"has_withdraw"`
		BeforeBalance      int64        `db:"before_balance"`
		BeforeWaitBalance  int64        `db:"before_wait_balance"`
		BeforeCny          int64        `db:"before_cny"`
		BeforeNotBilled    int64        `db:"before_not_billed"`
		BeforeBilled       int64        `db:"before_billed"`
		BeforeHasBilled    int64        `db:"before_has_billed"`
		BeforeWithdraw     int64        `db:"before_withdraw"`
		BeforeWaitWithdraw int64        `db:"before_wait_withdraw"`
		BeforeNotWithdraw  int64        `db:"before_not_withdraw"`
		BeforeHasWithdraw  int64        `db:"before_has_withdraw"`
		Remark             string       `db:"remark"`
		CreateAt           time.Time    `db:"create_at"`
		DeleteAt           sql.NullTime `db:"delete_at"`
	}
)

func newWalletRecordModel(conn sqlx.SqlConn) *defaultWalletRecordModel {
	return &defaultWalletRecordModel{
		conn:  conn,
		table: "`wallet_record`",
	}
}

func (m *defaultWalletRecordModel) withSession(session sqlx.Session) *defaultWalletRecordModel {
	return &defaultWalletRecordModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`wallet_record`",
	}
}

func (m *defaultWalletRecordModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultWalletRecordModel) FindOne(ctx context.Context, id int64) (*WalletRecord, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", walletRecordRows, m.table)
	var resp WalletRecord
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

func (m *defaultWalletRecordModel) Insert(ctx context.Context, data *WalletRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, walletRecordRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.UserId, data.Type, data.FundingId, data.Reason, data.Balance, data.WaitBalance, data.Cny, data.NotBilled, data.Billed, data.HasBilled, data.Withdraw, data.WaitWithdraw, data.NotWithdraw, data.HasWithdraw, data.BeforeBalance, data.BeforeWaitBalance, data.BeforeCny, data.BeforeNotBilled, data.BeforeBilled, data.BeforeHasBilled, data.BeforeWithdraw, data.BeforeWaitWithdraw, data.BeforeNotWithdraw, data.BeforeHasWithdraw, data.Remark, data.DeleteAt)
	return ret, err
}

func (m *defaultWalletRecordModel) Update(ctx context.Context, data *WalletRecord) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, walletRecordRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.UserId, data.Type, data.FundingId, data.Reason, data.Balance, data.WaitBalance, data.Cny, data.NotBilled, data.Billed, data.HasBilled, data.Withdraw, data.WaitWithdraw, data.NotWithdraw, data.HasWithdraw, data.BeforeBalance, data.BeforeWaitBalance, data.BeforeCny, data.BeforeNotBilled, data.BeforeBilled, data.BeforeHasBilled, data.BeforeWithdraw, data.BeforeWaitWithdraw, data.BeforeNotWithdraw, data.BeforeHasWithdraw, data.Remark, data.DeleteAt, data.Id)
	return err
}

func (m *defaultWalletRecordModel) tableName() string {
	return m.table
}
