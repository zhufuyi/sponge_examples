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
	userCommentCachePrefixKey = "userComment:"
	// UserCommentExpireTime expire time
	UserCommentExpireTime = 10 * time.Minute
)

var _ UserCommentCache = (*userCommentCache)(nil)

// UserCommentCache cache interface
type UserCommentCache interface {
	Set(ctx context.Context, id uint64, data *model.UserComment, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserComment, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserComment, error)
	MultiSet(ctx context.Context, data []*model.UserComment, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userCommentCache define a cache struct
type userCommentCache struct {
	cache cache.Cache
}

// NewUserCommentCache new a cache
func NewUserCommentCache(cacheType *model.CacheType) UserCommentCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserComment{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserComment{}
		})
	}

	return &userCommentCache{
		cache: c,
	}
}

// GetUserCommentCacheKey cache key
func (c *userCommentCache) GetUserCommentCacheKey(id uint64) string {
	return userCommentCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userCommentCache) Set(ctx context.Context, id uint64, data *model.UserComment, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserCommentCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userCommentCache) Get(ctx context.Context, id uint64) (*model.UserComment, error) {
	var data *model.UserComment
	cacheKey := c.GetUserCommentCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userCommentCache) MultiSet(ctx context.Context, data []*model.UserComment, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserCommentCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userCommentCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserComment, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserCommentCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserComment)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserComment)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserCommentCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userCommentCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCommentCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userCommentCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCommentCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
