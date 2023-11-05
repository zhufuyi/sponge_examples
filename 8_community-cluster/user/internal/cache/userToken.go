package cache

import (
	"context"
	"fmt"
	"strings"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
)

const (
	// cache prefix key, must end with a colon
	userTokenCachePrefixKey = "user:token:"
	// UserTokenExpireTime expire time
	UserTokenExpireTime = 24 * time.Hour
)

var _ UserTokenCache = (*userTokenCache)(nil)

// UserTokenCache cache interface
type UserTokenCache interface {
	Set(ctx context.Context, id uint64, token string, duration time.Duration) error
	Get(ctx context.Context, id uint64) (string, error)
	Del(ctx context.Context, id uint64) error
	Exist(ctx context.Context, id uint64) (int64, error)
}

type userTokenCache struct {
	cache cache.Cache
}

// NewUserTokenCache new cache
func NewUserTokenCache(cacheType *model.CacheType) UserTokenCache {
	newObject := func() interface{} {
		return ""
	}
	cachePrefix := ""
	jsonEncoding := encoding.JSONEncoding{}

	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, newObject)
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, newObject)
	}

	return &userTokenCache{
		cache: c,
	}
}

// Set cache
func (c *userTokenCache) Set(ctx context.Context, id uint64, token string, duration time.Duration) error {
	cacheKey := GetUserTokenCacheKey(id)
	return c.cache.Set(ctx, cacheKey, &token, duration)
}

// Get cache value
func (c *userTokenCache) Get(ctx context.Context, id uint64) (string, error) {
	var token string
	cacheKey := GetUserTokenCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &token)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Del delete cache
func (c *userTokenCache) Del(ctx context.Context, id uint64) error {
	cacheKey := GetUserTokenCacheKey(id)
	return c.cache.Del(ctx, cacheKey)
}

// Exist check key
func (c *userTokenCache) Exist(ctx context.Context, id uint64) (int64, error) {
	cacheKey := GetUserTokenCacheKey(id)
	return model.GetRedisCli().Exists(ctx, cacheKey).Result()
}

// GetUserTokenCacheKey cache key
func GetUserTokenCacheKey(id uint64) string {
	return fmt.Sprintf("%s%v", userTokenCachePrefixKey, id)
}
