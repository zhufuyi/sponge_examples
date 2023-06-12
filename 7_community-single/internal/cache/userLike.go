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
	userLikeCachePrefixKey = "userLike:"
	// UserLikeExpireTime expire time
	UserLikeExpireTime = 10 * time.Minute
)

var _ UserLikeCache = (*userLikeCache)(nil)

// UserLikeCache cache interface
type UserLikeCache interface {
	Set(ctx context.Context, id uint64, data *model.UserLike, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserLike, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserLike, error)
	MultiSet(ctx context.Context, data []*model.UserLike, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userLikeCache define a cache struct
type userLikeCache struct {
	cache cache.Cache
}

// NewUserLikeCache new a cache
func NewUserLikeCache(cacheType *model.CacheType) UserLikeCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserLike{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserLike{}
		})
	}

	return &userLikeCache{
		cache: c,
	}
}

// GetUserLikeCacheKey cache key
func (c *userLikeCache) GetUserLikeCacheKey(id uint64) string {
	return userLikeCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userLikeCache) Set(ctx context.Context, id uint64, data *model.UserLike, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserLikeCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userLikeCache) Get(ctx context.Context, id uint64) (*model.UserLike, error) {
	var data *model.UserLike
	cacheKey := c.GetUserLikeCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userLikeCache) MultiSet(ctx context.Context, data []*model.UserLike, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserLikeCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userLikeCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserLike, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserLikeCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserLike)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserLike)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserLikeCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userLikeCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserLikeCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userLikeCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserLikeCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
