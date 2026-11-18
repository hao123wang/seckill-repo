package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

type Options struct {
	addr         string // 地址，IP:PORT
	password     string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type Option func(*Options)

func WithAddr(addr string) Option {
	return func(o *Options) {
		o.addr = addr
	}
}

func WithPassword(pwd string) Option {
	return func(o *Options) {
		o.password = pwd
	}
}

func WithReadTimeout(readTimeout time.Duration) Option {
	return func(o *Options) {
		o.readTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(o *Options) {
		o.writeTimeout = writeTimeout
	}
}

// initRedis 初始化db连接
func Init(opts ...Option) {
	options := &Options{
		addr:         "127.0.0.1:6379",
		password:     "123456",
		readTimeout:  10 * time.Second,
		writeTimeout: 10 * time.Second,
	}
	for _, opt := range opts {
		opt(options)
	}

	newRedisConn(options)
}

func newRedisConn(options *Options) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         options.addr,
		Password:     options.password,
		ReadTimeout:  options.readTimeout,
		WriteTimeout: options.writeTimeout,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		rdb.Close()
		panic(fmt.Sprintf("Failed to ping redis, err: %v", err))
	}
}

func Close() {
	if rdb != nil {
		rdb.Close()
	}
}

// GetRedisCli 获取数据库连接
func GetRedisCli() *redis.Client {
	return rdb
}
