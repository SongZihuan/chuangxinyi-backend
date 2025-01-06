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
	announcementFieldNames          = builder.RawFieldNames(&Announcement{})
	announcementRows                = strings.Join(announcementFieldNames, ",")
	announcementRowsExpectAutoSet   = strings.Join(stringx.Remove(announcementFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	announcementRowsWithPlaceHolder = strings.Join(stringx.Remove(announcementFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	announcementModel interface {
		Insert(ctx context.Context, data *Announcement) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Announcement, error)
		FindOneById(ctx context.Context, id int64) (*Announcement, error)
		Update(ctx context.Context, data *Announcement) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAnnouncementModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Announcement struct {
		Id       int64        `db:"id"`
		Sort     int64        `db:"sort"`
		Title    string       `db:"title"`
		Content  string       `db:"content"`
		CreateAt time.Time    `db:"create_at"`
		StartAt  time.Time    `db:"start_at"`
		StopAt   time.Time    `db:"stop_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
	}
)

func newAnnouncementModel(conn sqlx.SqlConn) *defaultAnnouncementModel {
	return &defaultAnnouncementModel{
		conn:  conn,
		table: "`announcement`",
	}
}

func (m *defaultAnnouncementModel) withSession(session sqlx.Session) *defaultAnnouncementModel {
	return &defaultAnnouncementModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`announcement`",
	}
}

func (m *defaultAnnouncementModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAnnouncementModel) FindOne(ctx context.Context, id int64) (*Announcement, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", announcementRows, m.table)
	var resp Announcement
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

func (m *defaultAnnouncementModel) FindOneById(ctx context.Context, id int64) (*Announcement, error) {
	var resp Announcement
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", announcementRows, m.table)
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

func (m *defaultAnnouncementModel) Insert(ctx context.Context, data *Announcement) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, announcementRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Sort, data.Title, data.Content, data.StartAt, data.StopAt, data.DeleteAt)
	return ret, err
}

func (m *defaultAnnouncementModel) Update(ctx context.Context, newData *Announcement) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, announcementRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Sort, newData.Title, newData.Content, newData.StartAt, newData.StopAt, newData.DeleteAt, newData.Id)
	return err
}

func (m *defaultAnnouncementModel) tableName() string {
	return m.table
}