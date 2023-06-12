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
	commentCachePrefixKey = "comment:"
	// CommentExpireTime expire time
	CommentExpireTime = 10 * time.Minute
)

var _ CommentCache = (*commentCache)(nil)

// CommentCache cache interface
type CommentCache interface {
	Set(ctx context.Context, id uint64, data *model.Comment, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Comment, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Comment, error)
	MultiSet(ctx context.Context, data []*model.Comment, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// commentCache define a cache struct
type commentCache struct {
	cache cache.Cache
}

// NewCommentCache new a cache
func NewCommentCache(cacheType *model.CacheType) CommentCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Comment{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Comment{}
		})
	}

	return &commentCache{
		cache: c,
	}
}

// GetCommentCacheKey cache key
func (c *commentCache) GetCommentCacheKey(id uint64) string {
	return commentCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *commentCache) Set(ctx context.Context, id uint64, data *model.Comment, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCommentCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *commentCache) Get(ctx context.Context, id uint64) (*model.Comment, error) {
	var data *model.Comment
	cacheKey := c.GetCommentCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *commentCache) MultiSet(ctx context.Context, data []*model.Comment, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCommentCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *commentCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Comment, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCommentCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Comment)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Comment)
	for _, id := range ids {
		val, ok := itemMap[c.GetCommentCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *commentCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *commentCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
