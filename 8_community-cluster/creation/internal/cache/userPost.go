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
	userPostCachePrefixKey = "userPost:"
	// UserPostExpireTime expire time
	UserPostExpireTime = 10 * time.Minute
)

var _ UserPostCache = (*userPostCache)(nil)

// UserPostCache cache interface
type UserPostCache interface {
	Set(ctx context.Context, id uint64, data *model.UserPost, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserPost, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserPost, error)
	MultiSet(ctx context.Context, data []*model.UserPost, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userPostCache define a cache struct
type userPostCache struct {
	cache cache.Cache
}

// NewUserPostCache new a cache
func NewUserPostCache(cacheType *model.CacheType) UserPostCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserPost{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserPost{}
		})
	}

	return &userPostCache{
		cache: c,
	}
}

// GetUserPostCacheKey cache key
func (c *userPostCache) GetUserPostCacheKey(id uint64) string {
	return userPostCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userPostCache) Set(ctx context.Context, id uint64, data *model.UserPost, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserPostCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userPostCache) Get(ctx context.Context, id uint64) (*model.UserPost, error) {
	var data *model.UserPost
	cacheKey := c.GetUserPostCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userPostCache) MultiSet(ctx context.Context, data []*model.UserPost, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserPostCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userPostCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserPost, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserPostCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserPost)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserPost)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserPostCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userPostCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserPostCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userPostCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserPostCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
