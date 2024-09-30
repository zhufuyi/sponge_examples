package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"eshop/user/internal/model"
)

const (
	// cache prefix key, must end with a colon
	userCachePrefixKey = "user:"
	// UserExpireTime expire time
	UserExpireTime = 5 * time.Minute
)

var _ UserCache = (*userCache)(nil)

// UserCache cache interface
type UserCache interface {
	Set(ctx context.Context, id uint64, data *model.User, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.User, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.User, error)
	MultiSet(ctx context.Context, data []*model.User, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userCache define a cache struct
type userCache struct {
	cache cache.Cache
}

// NewUserCache new a cache
func NewUserCache(cacheType *model.CacheType) UserCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.User{}
		})
		return &userCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.User{}
		})
		return &userCache{cache: c}
	}

	return nil // no cache
}

// GetUserCacheKey cache key
func (c *userCache) GetUserCacheKey(id uint64) string {
	return userCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userCache) Set(ctx context.Context, id uint64, data *model.User, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userCache) Get(ctx context.Context, id uint64) (*model.User, error) {
	var data *model.User
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userCache) MultiSet(ctx context.Context, data []*model.User, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.User, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.User)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.User)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
