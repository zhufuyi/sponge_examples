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
	emailVerifyCodeCachePrefixKey = "user:emailVerifyCode:"
	// EmailVerifyCodeExpireTime expire time
	EmailVerifyCodeExpireTime = 10 * time.Minute
)

var _ EmailVerifyCodeCache = (*emailVerifyCodeCache)(nil)

// EmailVerifyCodeCache cache interface
type EmailVerifyCodeCache interface {
	Set(ctx context.Context, email string, verifyCode string, duration time.Duration) error
	Get(ctx context.Context, email string) (string, error)
	Del(ctx context.Context, email string) error
}

type emailVerifyCodeCache struct {
	cache cache.Cache
}

// NewEmailVerifyCodeCache new cache
func NewEmailVerifyCodeCache(cacheType *model.CacheType) EmailVerifyCodeCache {
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

	return &emailVerifyCodeCache{
		cache: c,
	}
}

// GetEmailVerifyCodeCacheKey cache key
func (c *emailVerifyCodeCache) getCacheKey(email string) string {
	return fmt.Sprintf("%s%v", emailVerifyCodeCachePrefixKey, email)
}

// Set cache
func (c *emailVerifyCodeCache) Set(ctx context.Context, email string, verifyCode string, duration time.Duration) error {
	cacheKey := c.getCacheKey(email)
	return c.cache.Set(ctx, cacheKey, &verifyCode, duration)
}

// Get cache value
func (c *emailVerifyCodeCache) Get(ctx context.Context, email string) (string, error) {
	var verifyCode string
	cacheKey := c.getCacheKey(email)
	err := c.cache.Get(ctx, cacheKey, &verifyCode)
	if err != nil {
		return "", err
	}
	return verifyCode, nil
}

// Del delete cache
func (c *emailVerifyCodeCache) Del(ctx context.Context, email string) error {
	cacheKey := c.getCacheKey(email)
	return c.cache.Del(ctx, cacheKey)
}
