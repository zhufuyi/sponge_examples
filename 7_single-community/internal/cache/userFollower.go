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
	userFollowerCachePrefixKey = "userFollower:"

	// UserFollowerExpireTime expire time
	UserFollowerExpireTime = 10 * time.Minute
)

var _ UserFollowerCache = (*userFollowerCache)(nil)

// UserFollowerCache cache interface
type UserFollowerCache interface {
	Set(ctx context.Context, id uint64, data *model.UserFollower, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserFollower, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserFollower, error)
	MultiSet(ctx context.Context, data []*model.UserFollower, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userFollowerCache define a cache struct
type userFollowerCache struct {
	cache cache.Cache
}

// NewUserFollowerCache new a cache
func NewUserFollowerCache(cacheType *model.CacheType) UserFollowerCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserFollower{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserFollower{}
		})
	}

	return &userFollowerCache{
		cache: c,
	}
}

// GetUserFollowerCacheKey cache key
func (c *userFollowerCache) GetUserFollowerCacheKey(id uint64) string {
	return userFollowerCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userFollowerCache) Set(ctx context.Context, id uint64, data *model.UserFollower, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserFollowerCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userFollowerCache) Get(ctx context.Context, id uint64) (*model.UserFollower, error) {
	var data *model.UserFollower
	cacheKey := c.GetUserFollowerCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userFollowerCache) MultiSet(ctx context.Context, data []*model.UserFollower, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserFollowerCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userFollowerCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserFollower, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserFollowerCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserFollower)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserFollower)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserFollowerCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userFollowerCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserFollowerCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userFollowerCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserFollowerCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
