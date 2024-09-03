package goredis

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisCli *redis.Client

	// ErrCacheNotFound No hit cache
	ErrCacheNotFound = redis.Nil
)

// InitRedis connecting to redis
// dsn supported formats.
// (1) no password, no db: localhost:6379
// (2) with password and db: <user>:<pass>@localhost:6379/2
// (3) redis://default:123456@localhost:6379/0?max_retries=3
// for more parameters see the redis source code for the setupConnParams function
func InitRedis(dsn string) error {
	opt, err := getRedisOptions(dsn)
	if err != nil {
		return err
	}

	redisCli = redis.NewClient(opt)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second) //nolint
	err = redisCli.Ping(ctx).Err()

	return err
}

func getRedisOptions(dsn string) (*redis.Options, error) {
	dsn = strings.ReplaceAll(dsn, " ", "")
	if len(dsn) > 8 {
		if !strings.Contains(dsn[len(dsn)-3:], "/") {
			dsn += "/0" // use db 0 by default
		}

		if dsn[:8] != "redis://" && dsn[:9] != "rediss://" {
			dsn = "redis://" + dsn
		}
	}

	return redis.ParseURL(dsn)
}

// GetRedisCli get redis client
func GetRedisCli() *redis.Client {
	if redisCli == nil {
		panic("redis client not initialized, call InitRedis first")
	}

	return redisCli
}

// CloseRedis close redis
func CloseRedis() error {
	if redisCli == nil {
		return nil
	}

	err := redisCli.Close()
	if err != nil && errors.Is(err, redis.ErrClosed) {
		return err
	}

	return nil
}
