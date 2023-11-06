package cache

import (
	"context"
	"strings"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	// cache prefix key, must end with a colon
	courseCachePrefixKey = "course:"
	// CourseExpireTime expire time
	CourseExpireTime = 10 * time.Minute
)

var _ CourseCache = (*courseCache)(nil)

// CourseCache cache interface
type CourseCache interface {
	Set(ctx context.Context, id uint64, data *model.Course, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Course, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Course, error)
	MultiSet(ctx context.Context, data []*model.Course, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// courseCache define a cache struct
type courseCache struct {
	cache cache.Cache
}

// NewCourseCache new a cache
func NewCourseCache(cacheType *model.CacheType) CourseCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Course{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Course{}
		})
	}

	return &courseCache{
		cache: c,
	}
}

// GetCourseCacheKey cache key
func (c *courseCache) GetCourseCacheKey(id uint64) string {
	return courseCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *courseCache) Set(ctx context.Context, id uint64, data *model.Course, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCourseCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *courseCache) Get(ctx context.Context, id uint64) (*model.Course, error) {
	var data *model.Course
	cacheKey := c.GetCourseCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *courseCache) MultiSet(ctx context.Context, data []*model.Course, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCourseCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *courseCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Course, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCourseCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Course)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Course)
	for _, id := range ids {
		val, ok := itemMap[c.GetCourseCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *courseCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCourseCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *courseCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCourseCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
