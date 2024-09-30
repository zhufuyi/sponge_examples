package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"pay/internal/model"
)

func newPayCache() *gotest.Cache {
	record1 := &model.Pay{}
	record1.ID = 1
	record2 := &model.Pay{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewPayCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_payCache_Set(t *testing.T) {
	c := newPayCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pay)
	err := c.ICache.(PayCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(PayCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_payCache_Get(t *testing.T) {
	c := newPayCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pay)
	err := c.ICache.(PayCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PayCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(PayCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_payCache_MultiGet(t *testing.T) {
	c := newPayCache()
	defer c.Close()

	var testData []*model.Pay
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Pay))
	}

	err := c.ICache.(PayCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PayCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Pay))
	}
}

func Test_payCache_MultiSet(t *testing.T) {
	c := newPayCache()
	defer c.Close()

	var testData []*model.Pay
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Pay))
	}

	err := c.ICache.(PayCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_payCache_Del(t *testing.T) {
	c := newPayCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pay)
	err := c.ICache.(PayCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_payCache_SetCacheWithNotFound(t *testing.T) {
	c := newPayCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pay)
	err := c.ICache.(PayCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewPayCache(t *testing.T) {
	c := NewPayCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewPayCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewPayCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
