package cache

import (
	"context"
	"strings"
	"time"

	"community/internal/model"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	// cache prefix key, must end with a colon
	postLatestCachePrefixKey = "postLatest:"
	// PostLatestExpireTime expire time
	PostLatestExpireTime = 10 * time.Minute
)

var _ PostLatestCache = (*postLatestCache)(nil)

// PostLatestCache cache interface
type PostLatestCache interface {
	Set(ctx context.Context, id uint64, data *model.PostLatest, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.PostLatest, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.PostLatest, error)
	MultiSet(ctx context.Context, data []*model.PostLatest, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// postLatestCache define a cache struct
type postLatestCache struct {
	cache cache.Cache
}

// NewPostLatestCache new a cache
func NewPostLatestCache(cacheType *model.CacheType) PostLatestCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.PostLatest{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.PostLatest{}
		})
	}

	return &postLatestCache{
		cache: c,
	}
}

// GetPostLatestCacheKey cache key
func (c *postLatestCache) GetPostLatestCacheKey(id uint64) string {
	return postLatestCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *postLatestCache) Set(ctx context.Context, id uint64, data *model.PostLatest, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPostLatestCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *postLatestCache) Get(ctx context.Context, id uint64) (*model.PostLatest, error) {
	var data *model.PostLatest
	cacheKey := c.GetPostLatestCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *postLatestCache) MultiSet(ctx context.Context, data []*model.PostLatest, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPostLatestCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *postLatestCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.PostLatest, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPostLatestCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.PostLatest)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.PostLatest)
	for _, id := range ids {
		val, ok := itemMap[c.GetPostLatestCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *postLatestCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPostLatestCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *postLatestCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetPostLatestCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
