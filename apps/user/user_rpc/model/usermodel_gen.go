// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheUserIdPrefix    = "cache:user:id:"
	cacheUserNamePrefix  = "cache:user:name:"
	cacheUserPhonePrefix = "cache:user:phone:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByName(ctx context.Context, name string) (*User, error)
		FindOneByPhone(ctx context.Context, phone int64) (*User, error)
		Update(ctx context.Context, newData *User) error
		Delete(ctx context.Context, id int64) error
		FindUsersByName(ctx context.Context, name string) ([]*User, error)
		ListUserByUid(ctx context.Context, id int64, length int) ([]*User, error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`
		Email      string    `db:"email"`
		Password   string    `db:"password"`
		Phone      int64     `db:"phone"`
		Address    string    `db:"address"`
		Gender     int64     `db:"gender"`
		Avatar     string    `db:"avatar"`
		Status     string    `db:"status"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	userNameKey := fmt.Sprintf("%s%v", cacheUserNamePrefix, data.Name)
	userPhoneKey := fmt.Sprintf("%s%v", cacheUserPhonePrefix, data.Phone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userIdKey, userNameKey, userPhoneKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	userNameKey := fmt.Sprintf("%s%v", cacheUserNamePrefix, name)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByPhone(ctx context.Context, phone int64) (*User, error) {
	userPhoneKey := fmt.Sprintf("%s%v", cacheUserPhonePrefix, phone)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindUsersByName(ctx context.Context, name string) ([]*User, error) {
	var resp []*User
	query := fmt.Sprintf("select %s from %s where `name` like %%%s%% limit 100", userRows, m.table, name)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (m *defaultUserModel) ListUserByUid(ctx context.Context, id int64, length int) ([]*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` > %d limit %d", userRows, m.table, id, length)
	var resp []*User
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userNameKey := fmt.Sprintf("%s%v", cacheUserNamePrefix, data.Name)
	userPhoneKey := fmt.Sprintf("%s%v", cacheUserPhonePrefix, data.Phone)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Email, data.Password, data.Phone, data.Address, data.Gender, data.Avatar, data.Status)
	}, userIdKey, userNameKey, userPhoneKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userNameKey := fmt.Sprintf("%s%v", cacheUserNamePrefix, data.Name)
	userPhoneKey := fmt.Sprintf("%s%v", cacheUserPhonePrefix, data.Phone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Email, newData.Password, newData.Phone, newData.Address, newData.Gender, newData.Avatar, newData.Status, newData.Id)
	}, userIdKey, userNameKey, userPhoneKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
