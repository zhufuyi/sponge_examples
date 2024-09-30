package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"pay/internal/model"
)

const (
	// cache prefix key, must end with a colon
	payCachePrefixKey = "pay:"
	// PayExpireTime expire time
	PayExpireTime = 5 * time.Minute
)

var _ PayCache = (*payCache)(nil)

// PayCache cache interface
type PayCache interface {
	Set(ctx context.Context, id uint64, data *model.Pay, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Pay, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Pay, error)
	MultiSet(ctx context.Context, data []*model.Pay, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// payCache define a cache struct
type payCache struct {
	cache cache.Cache
}

// NewPayCache new a cache
func NewPayCache(cacheType *model.CacheType) PayCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Pay{}
		})
		return &payCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Pay{}
		})
		return &payCache{cache: c}
	}

	return nil // no cache
}

// GetPayCacheKey cache key
func (c *payCache) GetPayCacheKey(id uint64) string {
	return payCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *payCache) Set(ctx context.Context, id uint64, data *model.Pay, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPayCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *payCache) Get(ctx context.Context, id uint64) (*model.Pay, error) {
	var data *model.Pay
	cacheKey := c.GetPayCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *payCache) MultiSet(ctx context.Context, data []*model.Pay, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPayCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *payCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Pay, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPayCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Pay)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Pay)
	for _, id := range ids {
		val, ok := itemMap[c.GetPayCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *payCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPayCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *payCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetPayCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
