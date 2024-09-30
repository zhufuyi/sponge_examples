package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"coupon/internal/model"
)

const (
	// cache prefix key, must end with a colon
	couponCachePrefixKey = "coupon:"
	// CouponExpireTime expire time
	CouponExpireTime = 5 * time.Minute
)

var _ CouponCache = (*couponCache)(nil)

// CouponCache cache interface
type CouponCache interface {
	Set(ctx context.Context, id uint64, data *model.Coupon, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Coupon, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Coupon, error)
	MultiSet(ctx context.Context, data []*model.Coupon, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// couponCache define a cache struct
type couponCache struct {
	cache cache.Cache
}

// NewCouponCache new a cache
func NewCouponCache(cacheType *model.CacheType) CouponCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Coupon{}
		})
		return &couponCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Coupon{}
		})
		return &couponCache{cache: c}
	}

	return nil // no cache
}

// GetCouponCacheKey cache key
func (c *couponCache) GetCouponCacheKey(id uint64) string {
	return couponCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *couponCache) Set(ctx context.Context, id uint64, data *model.Coupon, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCouponCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *couponCache) Get(ctx context.Context, id uint64) (*model.Coupon, error) {
	var data *model.Coupon
	cacheKey := c.GetCouponCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *couponCache) MultiSet(ctx context.Context, data []*model.Coupon, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCouponCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *couponCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Coupon, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCouponCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Coupon)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Coupon)
	for _, id := range ids {
		val, ok := itemMap[c.GetCouponCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *couponCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCouponCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *couponCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCouponCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
