package data

import (
	"fmt"

	"github.com/hao123wang/seckill-repo/packages/cache"
	"github.com/hao123wang/seckill-repo/packages/gormcli"
	"github.com/hao123wang/seckill-repo/user-srv/internal/biz"
	"github.com/hao123wang/seckill-repo/user-srv/internal/conf"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data holds the long-lived storage clients shared by repos.
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData opens the database client and returns it with a cleanup function.
func NewData(c *conf.Data) (*Data, error) {
	dbCfg := c.GetDatabase()
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.User, dbCfg.Password, dbCfg.Addr, dbCfg.Dbname)
	gormcli.Init(
		gormcli.WithSource(source),
		gormcli.WithMaxIdleConn(int(dbCfg.MaxIdleConn)),
		gormcli.WithMaxOpenConn(int(dbCfg.MaxOpenConn)),
	)

	rdbCfg := c.GetRedis()
	cache.Init(
		cache.WithAddr(rdbCfg.Addr),
		cache.WithPassword(rdbCfg.Password),
		cache.WithReadTimeout(rdbCfg.ReadTimeout.AsDuration()),
	)

	db := gormcli.GetDB()
	rdb := cache.GetRedisCli()

	if err := db.AutoMigrate(&biz.User{}); err != nil {
		return nil, err
	}
	return &Data{db: db, rdb: rdb}, nil
}
