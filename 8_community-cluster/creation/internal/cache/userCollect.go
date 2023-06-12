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
	userCollectCachePrefixKey = "userCollect:"
	// UserCollectExpireTime expire time
	UserCollectExpireTime = 10 * time.Minute
)

var _ UserCollectCache = (*userCollectCache)(nil)

// UserCollectCache cache interface
type UserCollectCache interface {
	Set(ctx context.Context, id uint64, data *model.UserCollect, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserCollect, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserCollect, error)
	MultiSet(ctx context.Context, data []*model.UserCollect, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userCollectCache define a cache struct
type userCollectCache struct {
	cache cache.Cache
}

// NewUserCollectCache new a cache
func NewUserCollectCache(cacheType *model.CacheType) UserCollectCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserCollect{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserCollect{}
		})
	}

	return &userCollectCache{
		cache: c,
	}
}

// GetUserCollectCacheKey cache key
func (c *userCollectCache) GetUserCollectCacheKey(id uint64) string {
	return userCollectCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userCollectCache) Set(ctx context.Context, id uint64, data *model.UserCollect, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserCollectCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userCollectCache) Get(ctx context.Context, id uint64) (*model.UserCollect, error) {
	var data *model.UserCollect
	cacheKey := c.GetUserCollectCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userCollectCache) MultiSet(ctx context.Context, data []*model.UserCollect, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserCollectCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userCollectCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserCollect, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserCollectCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserCollect)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserCollect)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserCollectCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userCollectCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCollectCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userCollectCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCollectCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
