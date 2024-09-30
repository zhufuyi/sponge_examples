package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"eshop/product/internal/model"
)

func newProductCache() *gotest.Cache {
	record1 := &model.Product{}
	record1.ID = 1
	record2 := &model.Product{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewProductCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_productCache_Set(t *testing.T) {
	c := newProductCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Product)
	err := c.ICache.(ProductCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(ProductCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_productCache_Get(t *testing.T) {
	c := newProductCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Product)
	err := c.ICache.(ProductCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(ProductCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_productCache_MultiGet(t *testing.T) {
	c := newProductCache()
	defer c.Close()

	var testData []*model.Product
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Product))
	}

	err := c.ICache.(ProductCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Product))
	}
}

func Test_productCache_MultiSet(t *testing.T) {
	c := newProductCache()
	defer c.Close()

	var testData []*model.Product
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Product))
	}

	err := c.ICache.(ProductCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productCache_Del(t *testing.T) {
	c := newProductCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Product)
	err := c.ICache.(ProductCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productCache_SetCacheWithNotFound(t *testing.T) {
	c := newProductCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Product)
	err := c.ICache.(ProductCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewProductCache(t *testing.T) {
	c := NewProductCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewProductCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewProductCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
