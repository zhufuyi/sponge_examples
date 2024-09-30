package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"product/internal/model"
)

const (
	// cache prefix key, must end with a colon
	productCachePrefixKey = "product:"
	// ProductExpireTime expire time
	ProductExpireTime = 5 * time.Minute
)

var _ ProductCache = (*productCache)(nil)

// ProductCache cache interface
type ProductCache interface {
	Set(ctx context.Context, id uint64, data *model.Product, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Product, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Product, error)
	MultiSet(ctx context.Context, data []*model.Product, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// productCache define a cache struct
type productCache struct {
	cache cache.Cache
}

// NewProductCache new a cache
func NewProductCache(cacheType *model.CacheType) ProductCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Product{}
		})
		return &productCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Product{}
		})
		return &productCache{cache: c}
	}

	return nil // no cache
}

// GetProductCacheKey cache key
func (c *productCache) GetProductCacheKey(id uint64) string {
	return productCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *productCache) Set(ctx context.Context, id uint64, data *model.Product, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetProductCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *productCache) Get(ctx context.Context, id uint64) (*model.Product, error) {
	var data *model.Product
	cacheKey := c.GetProductCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *productCache) MultiSet(ctx context.Context, data []*model.Product, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetProductCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *productCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Product, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetProductCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Product)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Product)
	for _, id := range ids {
		val, ok := itemMap[c.GetProductCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *productCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *productCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
