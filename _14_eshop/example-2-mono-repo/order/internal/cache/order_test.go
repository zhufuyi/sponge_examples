package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"eshop/order/internal/model"
)

func newOrderCache() *gotest.Cache {
	record1 := &model.OrderRecord{}
	record1.ID = 1
	record2 := &model.OrderRecord{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewOrderCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_orderCache_Set(t *testing.T) {
	c := newOrderCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.OrderRecord)
	err := c.ICache.(OrderCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(OrderCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_orderCache_Get(t *testing.T) {
	c := newOrderCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.OrderRecord)
	err := c.ICache.(OrderCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(OrderCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(OrderCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_orderCache_MultiGet(t *testing.T) {
	c := newOrderCache()
	defer c.Close()

	var testData []*model.OrderRecord
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.OrderRecord))
	}

	err := c.ICache.(OrderCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(OrderCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.OrderRecord))
	}
}

func Test_orderCache_MultiSet(t *testing.T) {
	c := newOrderCache()
	defer c.Close()

	var testData []*model.OrderRecord
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.OrderRecord))
	}

	err := c.ICache.(OrderCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_orderCache_Del(t *testing.T) {
	c := newOrderCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.OrderRecord)
	err := c.ICache.(OrderCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_orderCache_SetCacheWithNotFound(t *testing.T) {
	c := newOrderCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.OrderRecord)
	err := c.ICache.(OrderCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewOrderCache(t *testing.T) {
	c := NewOrderCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewOrderCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewOrderCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
