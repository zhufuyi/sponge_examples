package cache

import (
	"context"
	"strings"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	// cache prefix key, must end with a colon
	postHotCachePrefixKey = "postHot:"
	// PostHotExpireTime expire time
	PostHotExpireTime = 10 * time.Minute
)

var _ PostHotCache = (*postHotCache)(nil)

// PostHotCache cache interface
type PostHotCache interface {
	Set(ctx context.Context, id uint64, data *model.PostHot, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.PostHot, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.PostHot, error)
	MultiSet(ctx context.Context, data []*model.PostHot, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// postHotCache define a cache struct
type postHotCache struct {
	cache cache.Cache
}

// NewPostHotCache new a cache
func NewPostHotCache(cacheType *model.CacheType) PostHotCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.PostHot{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.PostHot{}
		})
	}

	return &postHotCache{
		cache: c,
	}
}

// GetPostHotCacheKey cache key
func (c *postHotCache) GetPostHotCacheKey(id uint64) string {
	return postHotCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *postHotCache) Set(ctx context.Context, id uint64, data *model.PostHot, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPostHotCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *postHotCache) Get(ctx context.Context, id uint64) (*model.PostHot, error) {
	var data *model.PostHot
	cacheKey := c.GetPostHotCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *postHotCache) MultiSet(ctx context.Context, data []*model.PostHot, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPostHotCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *postHotCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.PostHot, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPostHotCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.PostHot)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.PostHot)
	for _, id := range ids {
		val, ok := itemMap[c.GetPostHotCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *postHotCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPostHotCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *postHotCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetPostHotCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
