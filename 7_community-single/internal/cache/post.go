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
	postCachePrefixKey = "post:"
	// PostExpireTime expire time
	PostExpireTime = 10 * time.Minute
)

var _ PostCache = (*postCache)(nil)

// PostCache cache interface
type PostCache interface {
	Set(ctx context.Context, id uint64, data *model.Post, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Post, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Post, error)
	MultiSet(ctx context.Context, data []*model.Post, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// postCache define a cache struct
type postCache struct {
	cache cache.Cache
}

// NewPostCache new a cache
func NewPostCache(cacheType *model.CacheType) PostCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Post{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Post{}
		})
	}

	return &postCache{
		cache: c,
	}
}

// GetPostCacheKey cache key
func (c *postCache) GetPostCacheKey(id uint64) string {
	return postCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *postCache) Set(ctx context.Context, id uint64, data *model.Post, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPostCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *postCache) Get(ctx context.Context, id uint64) (*model.Post, error) {
	var data *model.Post
	cacheKey := c.GetPostCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *postCache) MultiSet(ctx context.Context, data []*model.Post, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPostCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *postCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Post, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPostCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Post)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Post)
	for _, id := range ids {
		val, ok := itemMap[c.GetPostCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *postCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPostCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *postCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetPostCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
