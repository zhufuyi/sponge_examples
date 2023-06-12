package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newPostHotCache() *gotest.Cache {
	record1 := &model.PostHot{}
	record1.ID = 1
	record2 := &model.PostHot{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewPostHotCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_postHotCache_Set(t *testing.T) {
	c := newPostHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostHot)
	err := c.ICache.(PostHotCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(PostHotCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_postHotCache_Get(t *testing.T) {
	c := newPostHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostHot)
	err := c.ICache.(PostHotCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PostHotCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(PostHotCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_postHotCache_MultiGet(t *testing.T) {
	c := newPostHotCache()
	defer c.Close()

	var testData []*model.PostHot
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.PostHot))
	}

	err := c.ICache.(PostHotCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PostHotCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.PostHot))
	}
}

func Test_postHotCache_MultiSet(t *testing.T) {
	c := newPostHotCache()
	defer c.Close()

	var testData []*model.PostHot
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.PostHot))
	}

	err := c.ICache.(PostHotCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_postHotCache_Del(t *testing.T) {
	c := newPostHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostHot)
	err := c.ICache.(PostHotCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_postHotCache_SetCacheWithNotFound(t *testing.T) {
	c := newPostHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostHot)
	err := c.ICache.(PostHotCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewPostHotCache(t *testing.T) {
	c := NewPostHotCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
