package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"eshop/order/internal/model"
)

const (
	// cache prefix key, must end with a colon
	orderCachePrefixKey = "order:"
	// OrderExpireTime expire time
	OrderExpireTime = 5 * time.Minute
)

var _ OrderCache = (*orderCache)(nil)

// OrderCache cache interface
type OrderCache interface {
	Set(ctx context.Context, id uint64, data *model.OrderRecord, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.OrderRecord, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.OrderRecord, error)
	MultiSet(ctx context.Context, data []*model.OrderRecord, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// orderCache define a cache struct
type orderCache struct {
	cache cache.Cache
}

// NewOrderCache new a cache
func NewOrderCache(cacheType *model.CacheType) OrderCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.OrderRecord{}
		})
		return &orderCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.OrderRecord{}
		})
		return &orderCache{cache: c}
	}

	return nil // no cache
}

// GetOrderCacheKey cache key
func (c *orderCache) GetOrderCacheKey(id uint64) string {
	return orderCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *orderCache) Set(ctx context.Context, id uint64, data *model.OrderRecord, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetOrderCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *orderCache) Get(ctx context.Context, id uint64) (*model.OrderRecord, error) {
	var data *model.OrderRecord
	cacheKey := c.GetOrderCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *orderCache) MultiSet(ctx context.Context, data []*model.OrderRecord, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetOrderCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *orderCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.OrderRecord, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetOrderCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.OrderRecord)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.OrderRecord)
	for _, id := range ids {
		val, ok := itemMap[c.GetOrderCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *orderCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetOrderCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *orderCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetOrderCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
