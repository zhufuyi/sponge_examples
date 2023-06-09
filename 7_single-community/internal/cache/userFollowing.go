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
	userFollowingCachePrefixKey = "userFollowing:"

	// UserFollowingExpireTime expire time
	UserFollowingExpireTime = 10 * time.Minute
)

var _ UserFollowingCache = (*userFollowingCache)(nil)

// UserFollowingCache cache interface
type UserFollowingCache interface {
	Set(ctx context.Context, id uint64, data *model.UserFollowing, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserFollowing, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserFollowing, error)
	MultiSet(ctx context.Context, data []*model.UserFollowing, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userFollowingCache define a cache struct
type userFollowingCache struct {
	cache cache.Cache
}

// NewUserFollowingCache new a cache
func NewUserFollowingCache(cacheType *model.CacheType) UserFollowingCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserFollowing{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserFollowing{}
		})
	}

	return &userFollowingCache{
		cache: c,
	}
}

// GetUserFollowingCacheKey cache key
func (c *userFollowingCache) GetUserFollowingCacheKey(id uint64) string {
	return userFollowingCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userFollowingCache) Set(ctx context.Context, id uint64, data *model.UserFollowing, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserFollowingCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userFollowingCache) Get(ctx context.Context, id uint64) (*model.UserFollowing, error) {
	var data *model.UserFollowing
	cacheKey := c.GetUserFollowingCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userFollowingCache) MultiSet(ctx context.Context, data []*model.UserFollowing, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserFollowingCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userFollowingCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserFollowing, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserFollowingCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserFollowing)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserFollowing)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserFollowingCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userFollowingCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserFollowingCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userFollowingCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserFollowingCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
