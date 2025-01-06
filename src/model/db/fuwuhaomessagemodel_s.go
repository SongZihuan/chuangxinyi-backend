package db

import (
	"context"
	"fmt"
	"gitee.com/wuntsong-auth/backend/src/where"
	"github.com/wuntsong-org/go-zero-plus/core/stores/sqlc"
)

type (
	fuwuhaoMessageModelSelf interface {
		GetList(ctx context.Context, openID string, page int64, pageSize int64, startTime, endTime int64, timeType int64, senderID int64) ([]FuwuhaoMessage, error)
		GetCount(ctx context.Context, openID string, startTime, endTime int64, timeType int64, senderID int64) (int64, error)
	}
)

func (m *defaultFuwuhaoMessageModel) GetList(ctx context.Context, openID string, page int64, pageSize int64, startTime, endTime int64, timeType int64, senderID int64) ([]FuwuhaoMessage, error) {
	var resp []FuwuhaoMessage
	var err error

	cond := where.NewCond(m.table, fuwuhaoMessageFieldNames).StringEQ("open_id", openID).WebIDWithCenter(senderID, "sender_id").TimeBetweenWithTimeMap(timeType, timeMap, startTime, endTime)
	query := fmt.Sprintf("select %s from %s where %s order by %s %s", fuwuhaoMessageRows, m.table, cond, cond.OrderBy(), where.NewPage(page, pageSize))

	err = m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return []FuwuhaoMessage{}, nil
	default:
		return nil, err
	}
}

func (m *defaultFuwuhaoMessageModel) GetCount(ctx context.Context, openID string, startTime, endTime int64, timeType int64, senderID int64) (int64, error) {
	var err error
	var resp OneInt

	cond := where.NewCond(m.table, fuwuhaoMessageFieldNames).StringEQ("open_id", openID).WebIDWithCenter(senderID, "sender_id").TimeBetweenWithTimeMap(timeType, timeMap, startTime, endTime)
	query := fmt.Sprintf("select count(id) as res from %s where %s", m.table, cond)

	err = m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp.Res, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}