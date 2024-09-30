package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"stock/internal/model"
)

const (
	// cache prefix key, must end with a colon
	stockCachePrefixKey = "stock:"
	// StockExpireTime expire time
	StockExpireTime = 5 * time.Minute
)

var _ StockCache = (*stockCache)(nil)

// StockCache cache interface
type StockCache interface {
	Set(ctx context.Context, id uint64, data *model.Stock, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Stock, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Stock, error)
	MultiSet(ctx context.Context, data []*model.Stock, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// stockCache define a cache struct
type stockCache struct {
	cache cache.Cache
}

// NewStockCache new a cache
func NewStockCache(cacheType *model.CacheType) StockCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Stock{}
		})
		return &stockCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Stock{}
		})
		return &stockCache{cache: c}
	}

	return nil // no cache
}

// GetStockCacheKey cache key
func (c *stockCache) GetStockCacheKey(id uint64) string {
	return stockCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *stockCache) Set(ctx context.Context, id uint64, data *model.Stock, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetStockCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *stockCache) Get(ctx context.Context, id uint64) (*model.Stock, error) {
	var data *model.Stock
	cacheKey := c.GetStockCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *stockCache) MultiSet(ctx context.Context, data []*model.Stock, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetStockCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *stockCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Stock, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetStockCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Stock)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Stock)
	for _, id := range ids {
		val, ok := itemMap[c.GetStockCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *stockCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetStockCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *stockCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetStockCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
