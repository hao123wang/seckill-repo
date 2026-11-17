package gormcli

import (
	"context"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ctxTransactionKey struct {
}

var (
	db *gorm.DB
)

type Options struct {
	source                   string
	maxIdleConn              int
	maxOpenConn              int
	slowThresholdMillisecond int
}

type Option func(*Options)

func WithSlowThresholdMillisecond(slowThresholdMillisecond int) Option {
	return func(o *Options) {
		o.slowThresholdMillisecond = slowThresholdMillisecond
	}
}

func WithSource(source string) Option {
	return func(o *Options) {
		o.source = source
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(o *Options) {
		o.maxIdleConn = maxIdleConn
	}
}

func newOptions(opts ...Option) Options {
	options := Options{
		maxIdleConn:              10,
		maxOpenConn:              10,
		slowThresholdMillisecond: 0,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

func Init(opts ...Option) {
	options := &Options{
		source:                   "",
		maxIdleConn:              10,
		maxOpenConn:              20,
		slowThresholdMillisecond: 0,
	}
	for _, opt := range opts {
		opt(options)
	}

	newDB(options)
}

func newDB(options *Options) {
	var err error
	db, err = gorm.Open(mysql.Open(options.source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("fetch gormCli connection err:" + err.Error())
	}

	sqlDB.SetMaxIdleConns(options.maxIdleConn) // 设置最大空闲连接
	sqlDB.SetMaxOpenConns(options.maxOpenConn) // 设置最大打开的连接
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}

func Close() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			slog.Info("get sql db failed", "err", err)
		}
		sqlDB.Close()
	}
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error) error {
	db := GetDB().WithContext(ctx)

	return db.Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fc(txctx)
	})
}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	c := context.WithValue(ctx, ctxTransactionKey{}, tx)
	return c
}
