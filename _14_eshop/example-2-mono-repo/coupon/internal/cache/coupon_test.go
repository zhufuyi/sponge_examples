package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"eshop/coupon/internal/model"
)

func newCouponCache() *gotest.Cache {
	record1 := &model.Coupon{}
	record1.ID = 1
	record2 := &model.Coupon{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCouponCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_couponCache_Set(t *testing.T) {
	c := newCouponCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Coupon)
	err := c.ICache.(CouponCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CouponCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_couponCache_Get(t *testing.T) {
	c := newCouponCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Coupon)
	err := c.ICache.(CouponCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CouponCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CouponCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_couponCache_MultiGet(t *testing.T) {
	c := newCouponCache()
	defer c.Close()

	var testData []*model.Coupon
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Coupon))
	}

	err := c.ICache.(CouponCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CouponCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Coupon))
	}
}

func Test_couponCache_MultiSet(t *testing.T) {
	c := newCouponCache()
	defer c.Close()

	var testData []*model.Coupon
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Coupon))
	}

	err := c.ICache.(CouponCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_couponCache_Del(t *testing.T) {
	c := newCouponCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Coupon)
	err := c.ICache.(CouponCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_couponCache_SetCacheWithNotFound(t *testing.T) {
	c := newCouponCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Coupon)
	err := c.ICache.(CouponCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCouponCache(t *testing.T) {
	c := NewCouponCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewCouponCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewCouponCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
