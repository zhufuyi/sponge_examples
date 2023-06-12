package cache

import (
	"context"
	"strings"
	"time"

	"relation/internal/model"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	//  cache prefix key, must end with a colon
	relationNumCachePrefixKey = "relationNum:"

	// cache prefix key, must end with a colon
	relationNumUidCachePrefixKey = "relationNum:uid:"
	// RelationNumExpireTime expire time
	RelationNumExpireTime = 10 * time.Minute // nolint
)

var _ RelationNumCache = (*relationNumCache)(nil)

// RelationNumCache cache interface
type RelationNumCache interface {
	Set(ctx context.Context, id uint64, data *model.RelationNum, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.RelationNum, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.RelationNum, error)
	MultiSet(ctx context.Context, data []*model.RelationNum, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error

	SetByUid(ctx context.Context, userID uint64, data *model.RelationNum, duration time.Duration) error
	GetByUid(ctx context.Context, userID uint64) (*model.RelationNum, error)
	DelByUid(ctx context.Context, userID uint64) error
	SetByUidCacheWithNotFound(ctx context.Context, userID uint64) error
}

// relationNumCache define a cache struct
type relationNumCache struct {
	cache cache.Cache
}

// NewRelationNumCache new a cache
func NewRelationNumCache(cacheType *model.CacheType) RelationNumCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	var c cache.Cache
	if strings.ToLower(cacheType.CType) == "redis" {
		c = cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.RelationNum{}
		})
	} else {
		c = cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.RelationNum{}
		})
	}

	return &relationNumCache{
		cache: c,
	}
}

// GetRelationNumCacheKey cache key
func (c *relationNumCache) GetRelationNumCacheKey(id uint64) string {
	return relationNumCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *relationNumCache) Set(ctx context.Context, id uint64, data *model.RelationNum, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetRelationNumCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *relationNumCache) Get(ctx context.Context, id uint64) (*model.RelationNum, error) {
	var data *model.RelationNum
	cacheKey := c.GetRelationNumCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *relationNumCache) MultiSet(ctx context.Context, data []*model.RelationNum, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetRelationNumCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *relationNumCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.RelationNum, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetRelationNumCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.RelationNum)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.RelationNum)
	for _, id := range ids {
		val, ok := itemMap[c.GetRelationNumCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *relationNumCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetRelationNumCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *relationNumCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetRelationNumCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// cache key
func (c *relationNumCache) getByUidCacheKey(userID uint64) string {
	return relationNumUidCachePrefixKey + utils.Uint64ToStr(userID)
}

// Set cache
func (c *relationNumCache) SetByUid(ctx context.Context, userID uint64, data *model.RelationNum, duration time.Duration) error {
	cacheKey := c.getByUidCacheKey(userID)
	return c.cache.Set(ctx, cacheKey, data, duration)
}

// Get cache
func (c *relationNumCache) GetByUid(ctx context.Context, userID uint64) (*model.RelationNum, error) {
	var data *model.RelationNum
	cacheKey := c.getByUidCacheKey(userID)
	err := c.cache.Get(ctx, cacheKey, data)
	return data, err
}

// Del delete cache
func (c *relationNumCache) DelByUid(ctx context.Context, userID uint64) error {
	cacheKey := c.getByUidCacheKey(userID)
	return c.cache.Del(ctx, cacheKey)
}

// SetByUidCacheWithNotFound set empty cache
func (c *relationNumCache) SetByUidCacheWithNotFound(ctx context.Context, userID uint64) error {
	cacheKey := c.getByUidCacheKey(userID)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
