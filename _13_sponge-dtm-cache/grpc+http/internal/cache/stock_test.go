package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"stock/internal/model"
)

func newStockCache() *gotest.Cache {
	record1 := &model.Stock{}
	record1.ID = 1
	record2 := &model.Stock{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewStockCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_stockCache_Set(t *testing.T) {
	c := newStockCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Stock)
	err := c.ICache.(StockCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(StockCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_stockCache_Get(t *testing.T) {
	c := newStockCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Stock)
	err := c.ICache.(StockCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(StockCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(StockCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_stockCache_MultiGet(t *testing.T) {
	c := newStockCache()
	defer c.Close()

	var testData []*model.Stock
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Stock))
	}

	err := c.ICache.(StockCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(StockCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Stock))
	}
}

func Test_stockCache_MultiSet(t *testing.T) {
	c := newStockCache()
	defer c.Close()

	var testData []*model.Stock
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Stock))
	}

	err := c.ICache.(StockCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_stockCache_Del(t *testing.T) {
	c := newStockCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Stock)
	err := c.ICache.(StockCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_stockCache_SetCacheWithNotFound(t *testing.T) {
	c := newStockCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Stock)
	err := c.ICache.(StockCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewStockCache(t *testing.T) {
	c := NewStockCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewStockCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewStockCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
