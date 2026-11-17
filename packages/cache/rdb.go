package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	client    = &Client{}
	redisConn *redis.Client
)

type Client struct {
}

type Options struct {
	addr         string // 地址，IP:PORT
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type Option func(*Options)

func WithAddr(addr string) Option {
	return func(o *Options) {
		o.addr = addr
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
		readTimeout:  10,
		writeTimeout: 10,
	}
	for _, opt := range opts {
		opt(options)
	}

	newRedisConn(options)
}

func newRedisConn(options *Options) {
	redisConn = redis.NewClient(&redis.Options{
		Addr:         options.addr,
		ReadTimeout:  options.readTimeout,
		WriteTimeout: options.writeTimeout,
	})
	if redisConn == nil {
		panic("failed to call redis.NewClient")
	}
	_, err := redisConn.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to ping redis")
	}
}

func (client *Client) Close() {
	if redisConn != nil {
		redisConn.Close()
	}
}

// GetRedisCli 获取数据库连接
func GetRedisCli() *Client {
	return client
}
