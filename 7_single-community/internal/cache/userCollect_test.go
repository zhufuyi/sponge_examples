package cache

import (
	"testing"
	"time"

	"community/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserCollectCache() *gotest.Cache {
	record1 := &model.UserCollect{}
	record1.ID = 1
	record2 := &model.UserCollect{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserCollectCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userCollectCache_Set(t *testing.T) {
	c := newUserCollectCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserCollect)
	err := c.ICache.(UserCollectCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserCollectCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userCollectCache_Get(t *testing.T) {
	c := newUserCollectCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserCollect)
	err := c.ICache.(UserCollectCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserCollectCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserCollectCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userCollectCache_MultiGet(t *testing.T) {
	c := newUserCollectCache()
	defer c.Close()

	var testData []*model.UserCollect
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserCollect))
	}

	err := c.ICache.(UserCollectCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserCollectCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserCollect))
	}
}

func Test_userCollectCache_MultiSet(t *testing.T) {
	c := newUserCollectCache()
	defer c.Close()

	var testData []*model.UserCollect
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserCollect))
	}

	err := c.ICache.(UserCollectCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userCollectCache_Del(t *testing.T) {
	c := newUserCollectCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserCollect)
	err := c.ICache.(UserCollectCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userCollectCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserCollectCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserCollect)
	err := c.ICache.(UserCollectCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserCollectCache(t *testing.T) {
	c := NewUserCollectCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
