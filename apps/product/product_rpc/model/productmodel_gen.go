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
	productFieldNames          = builder.RawFieldNames(&Product{})
	productRows                = strings.Join(productFieldNames, ",")
	productRowsExpectAutoSet   = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	productRowsWithPlaceHolder = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheProductIdPrefix = "cache:product:id:"
)

type (
	productModel interface {
		Insert(ctx context.Context, data *Product) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Product, error)
		Update(ctx context.Context, data *Product) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultProductModel struct {
		sqlc.CachedConn
		table string
	}

	Product struct {
		Id         uint64    `db:"id"`
		Name       string    `db:"name"`   // 产品名称
		Desc       string    `db:"desc"`   // 产品描述
		Stock      uint64    `db:"stock"`  // 产品库存
		Amount     uint64    `db:"amount"` // 产品金额
		Status     uint64    `db:"status"` // 产品状态
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newProductModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductModel {
	return &defaultProductModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`product`",
	}
}

func (m *defaultProductModel) Delete(ctx context.Context, id uint64) error {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productIdKey)
	return err
}

func (m *defaultProductModel) FindOne(ctx context.Context, id uint64) (*Product, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, id)
	var resp Product
	err := m.QueryRowCtx(ctx, &resp, productIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
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

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, productRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Desc, data.Stock, data.Amount, data.Status)
	}, productIdKey)
	return ret, err
}

func (m *defaultProductModel) Update(ctx context.Context, data *Product) error {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Desc, data.Stock, data.Amount, data.Status, data.Id)
	}, productIdKey)
	return err
}

func (m *defaultProductModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProductIdPrefix, primary)
}

func (m *defaultProductModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductModel) tableName() string {
	return m.table
}