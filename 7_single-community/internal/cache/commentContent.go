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
	commentContentCachePrefixKey = "commentContent:"
	// CommentContentExpireTime expire time
	CommentContentExpireTime = 10 * time.Minute
)

var _ CommentContentCache = (*commentContentCache)(nil)

// CommentContentCache cache interface
type CommentContentCache interface {
	Set(ctx context.Context, id uint64, data *model.CommentContent, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.CommentContent, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CommentContent, error)
	MultiSet(ctx context.Context, data []*model.CommentContent, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error

	MultiSetByCommentID(ctx context.Context, contents []*model.CommentContent, duration time.Duration) error
	MultiGetByCommentID(ctx context.Context, commentIDs []uint64) (map[uint64]*model.CommentContent, error)
}

// commentContentCache define a cache struct
type commentContentCache struct {
	cache cache.Cache
}

// NewCommentContentCache new a cache
func NewCommentContentCache(cacheType *model.CacheType) CommentContentCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.CommentContent{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.CommentContent{}
		})
	}

	return &commentContentCache{
		cache: c,
	}
}

// GetCommentContentCacheKey cache key
func (c *commentContentCache) GetCommentContentCacheKey(id uint64) string {
	return commentContentCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *commentContentCache) Set(ctx context.Context, id uint64, data *model.CommentContent, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCommentContentCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *commentContentCache) Get(ctx context.Context, id uint64) (*model.CommentContent, error) {
	var data *model.CommentContent
	cacheKey := c.GetCommentContentCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *commentContentCache) MultiSet(ctx context.Context, data []*model.CommentContent, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCommentContentCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiSetByCommentID multiple set cache by comment id
func (c *commentContentCache) MultiSetByCommentID(ctx context.Context, contents []*model.CommentContent, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range contents {
		cacheKey := c.GetCommentContentCacheKey(v.CommentID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *commentContentCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CommentContent, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCommentContentCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.CommentContent)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.CommentContent)
	for _, id := range ids {
		val, ok := itemMap[c.GetCommentContentCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// MultiGetByCommentID multiple get cache, return key in map is id value
func (c *commentContentCache) MultiGetByCommentID(ctx context.Context, commentIDs []uint64) (map[uint64]*model.CommentContent, error) {
	var keys []string
	for _, v := range commentIDs {
		cacheKey := c.GetCommentContentCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.CommentContent)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.CommentContent)
	for _, commentID := range commentIDs {
		val, ok := itemMap[c.GetCommentContentCacheKey(commentID)]
		if ok {
			retMap[commentID] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *commentContentCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentContentCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *commentContentCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCommentContentCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
