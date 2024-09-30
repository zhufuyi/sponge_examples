package data

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/zhufuyi/sponge/pkg/logger"

	"flashSale/internal/config"
)

var (
	// ErrCacheNotFound No hit cache
	ErrCacheNotFound = redis.Nil
)

var (
	redisCli  *redis.Client
	redisOnce sync.Once
)

// InitRedis connecting to redis
// dsn supported formats.
// (1) no password, no db: localhost:6379
// (2) with password and db: <user>:<pass>@localhost:6379/2
// (3) redis://default:123456@localhost:6379/0?max_retries=3
// for more parameters see the redis source code for the setupConnParams function
func InitRedis() error {
	opt, err := getRedisOptions()
	if err != nil {
		return err
	}

	redisCli = redis.NewClient(opt)

	if config.Get().App.EnableTrace {
		redisCli.AddHook(redisotel.TracingHook{})
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second) //nolint
	err = redisCli.Ping(ctx).Err()

	return nil
}

func getRedisOptions() (*redis.Options, error) {
	dsn := strings.ReplaceAll(config.Get().Redis.Dsn, " ", "")
	if len(dsn) > 8 {
		if !strings.Contains(dsn[len(dsn)-3:], "/") {
			dsn += "/0" // use db 0 by default
		}

		if dsn[:8] != "redis://" && dsn[:9] != "rediss://" {
			dsn = "redis://" + dsn
		}
	}

	redisOpts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	if config.Get().Redis.DialTimeout > 0 {
		redisOpts.DialTimeout = time.Duration(config.Get().Redis.DialTimeout) * time.Second
	}
	if config.Get().Redis.ReadTimeout > 0 {
		redisOpts.ReadTimeout = time.Duration(config.Get().Redis.ReadTimeout) * time.Second
	}
	if config.Get().Redis.WriteTimeout > 0 {
		redisOpts.WriteTimeout = time.Duration(config.Get().Redis.WriteTimeout) * time.Second
	}

	return redisOpts, nil
}

// GetRedisCli get redis client
func GetRedisCli() *redis.Client {
	if redisCli == nil {
		redisOnce.Do(func() {
			err := InitRedis()
			if err != nil {
				panic(fmt.Errorf("redis init failed: %v", err))
			}
		})
	}

	return redisCli
}

// CloseRedis close redis
func CloseRedis() error {
	if redisCli == nil {
		return nil
	}

	err := redisCli.Close()
	if err != nil && err.Error() != redis.ErrClosed.Error() {
		return err
	}

	return nil
}

// GetStockCacheKey get redis key for stock
func GetStockCacheKey(productID uint64) string {
	return "stock" + ":product_id:" + strconv.FormatUint(productID, 10)
}

// SetStock set redis key for stock
func SetStock(productID uint64, stock int) error {
	stockKey := GetStockCacheKey(productID)
	_, err := GetRedisCli().Set(context.Background(), stockKey, stock, 86400*time.Second).Result()
	return err
}

func SetDefaultStock() {
	var productID uint64 = 1
	var stock int = 5
	err := SetStock(productID, stock)
	if err != nil {
		logger.Errorf("set demo failed, err=%v", err)
		return
	}
	logger.Infof("set demo, key=%s, stock=%d", GetStockCacheKey(productID), stock)
}
