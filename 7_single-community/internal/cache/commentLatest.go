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
	commentLatestCachePrefixKey = "commentLatest:"
	// CommentLatestExpireTime expire time
	CommentLatestExpireTime = 10 * time.Minute
)

var _ CommentLatestCache = (*commentLatestCache)(nil)

// CommentLatestCache cache interface
type CommentLatestCache interface {
	Set(ctx context.Context, id uint64, data *model.CommentLatest, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.CommentLatest, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CommentLatest, error)
	MultiSet(ctx context.Context, data []*model.CommentLatest, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// commentLatestCache define a cache struct
type commentLatestCache struct {
	cache cache.Cache
}

// NewCommentLatestCache new a cache
func NewCommentLatestCache(cacheType *model.CacheType) CommentLatestCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.CommentLatest{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.CommentLatest{}
		})
	}

	return &commentLatestCache{
		cache: c,
	}
}

// GetCommentLatestCacheKey cache key
func (c *commentLatestCache) GetCommentLatestCacheKey(id uint64) string {
	return commentLatestCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *commentLatestCache) Set(ctx context.Context, id uint64, data *model.CommentLatest, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCommentLatestCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *commentLatestCache) Get(ctx context.Context, id uint64) (*model.CommentLatest, error) {
	var data *model.CommentLatest
	cacheKey := c.GetCommentLatestCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *commentLatestCache) MultiSet(ctx context.Context, data []*model.CommentLatest, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCommentLatestCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *commentLatestCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CommentLatest, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCommentLatestCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.CommentLatest)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.CommentLatest)
	for _, id := range ids {
		val, ok := itemMap[c.GetCommentLatestCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *commentLatestCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentLatestCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *commentLatestCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentLatestCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
