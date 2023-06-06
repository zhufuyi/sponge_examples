package cache

import (
	"context"
	"strings"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	// cache prefix key, must end with a colon
	teachCachePrefixKey = "teach:"
	// TeachExpireTime expire time
	TeachExpireTime = 10 * time.Minute
)

var _ TeachCache = (*teachCache)(nil)

// TeachCache cache interface
type TeachCache interface {
	Set(ctx context.Context, id uint64, data *model.Teach, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Teach, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Teach, error)
	MultiSet(ctx context.Context, data []*model.Teach, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// teachCache define a cache struct
type teachCache struct {
	cache cache.Cache
}

// NewTeachCache new a cache
func NewTeachCache(cacheType *model.CacheType) TeachCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Teach{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Teach{}
		})
	}

	return &teachCache{
		cache: c,
	}
}

// GetTeachCacheKey cache key
func (c *teachCache) GetTeachCacheKey(id uint64) string {
	return teachCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *teachCache) Set(ctx context.Context, id uint64, data *model.Teach, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetTeachCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *teachCache) Get(ctx context.Context, id uint64) (*model.Teach, error) {
	var data *model.Teach
	cacheKey := c.GetTeachCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *teachCache) MultiSet(ctx context.Context, data []*model.Teach, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetTeachCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *teachCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Teach, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetTeachCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Teach)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Teach)
	for _, id := range ids {
		val, ok := itemMap[c.GetTeachCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *teachCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetTeachCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *teachCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetTeachCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
