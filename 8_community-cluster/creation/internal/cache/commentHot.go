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
	commentHotCachePrefixKey = "commentHot:"
	// CommentHotExpireTime expire time
	CommentHotExpireTime = 10 * time.Minute
)

var _ CommentHotCache = (*commentHotCache)(nil)

// CommentHotCache cache interface
type CommentHotCache interface {
	Set(ctx context.Context, id uint64, data *model.CommentHot, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.CommentHot, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CommentHot, error)
	MultiSet(ctx context.Context, data []*model.CommentHot, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// commentHotCache define a cache struct
type commentHotCache struct {
	cache cache.Cache
}

// NewCommentHotCache new a cache
func NewCommentHotCache(cacheType *model.CacheType) CommentHotCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.CommentHot{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.CommentHot{}
		})
	}

	return &commentHotCache{
		cache: c,
	}
}

// GetCommentHotCacheKey cache key
func (c *commentHotCache) GetCommentHotCacheKey(id uint64) string {
	return commentHotCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *commentHotCache) Set(ctx context.Context, id uint64, data *model.CommentHot, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCommentHotCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *commentHotCache) Get(ctx context.Context, id uint64) (*model.CommentHot, error) {
	var data *model.CommentHot
	cacheKey := c.GetCommentHotCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *commentHotCache) MultiSet(ctx context.Context, data []*model.CommentHot, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCommentHotCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *commentHotCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CommentHot, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCommentHotCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.CommentHot)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.CommentHot)
	for _, id := range ids {
		val, ok := itemMap[c.GetCommentHotCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *commentHotCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentHotCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *commentHotCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentHotCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
