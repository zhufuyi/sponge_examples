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
	teacherCachePrefixKey = "teacher:"
	// TeacherExpireTime expire time
	TeacherExpireTime = 10 * time.Minute
)

var _ TeacherCache = (*teacherCache)(nil)

// TeacherCache cache interface
type TeacherCache interface {
	Set(ctx context.Context, id uint64, data *model.Teacher, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Teacher, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Teacher, error)
	MultiSet(ctx context.Context, data []*model.Teacher, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// teacherCache define a cache struct
type teacherCache struct {
	cache cache.Cache
}

// NewTeacherCache new a cache
func NewTeacherCache(cacheType *model.CacheType) TeacherCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Teacher{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Teacher{}
		})
	}

	return &teacherCache{
		cache: c,
	}
}

// GetTeacherCacheKey cache key
func (c *teacherCache) GetTeacherCacheKey(id uint64) string {
	return teacherCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *teacherCache) Set(ctx context.Context, id uint64, data *model.Teacher, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetTeacherCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *teacherCache) Get(ctx context.Context, id uint64) (*model.Teacher, error) {
	var data *model.Teacher
	cacheKey := c.GetTeacherCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *teacherCache) MultiSet(ctx context.Context, data []*model.Teacher, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetTeacherCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *teacherCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Teacher, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetTeacherCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Teacher)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Teacher)
	for _, id := range ids {
		val, ok := itemMap[c.GetTeacherCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *teacherCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetTeacherCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *teacherCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetTeacherCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
