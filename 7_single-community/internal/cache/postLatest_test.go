package cache

import (
	"testing"
	"time"

	"community/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newPostLatestCache() *gotest.Cache {
	record1 := &model.PostLatest{}
	record1.ID = 1
	record2 := &model.PostLatest{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewPostLatestCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_postLatestCache_Set(t *testing.T) {
	c := newPostLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostLatest)
	err := c.ICache.(PostLatestCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(PostLatestCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_postLatestCache_Get(t *testing.T) {
	c := newPostLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostLatest)
	err := c.ICache.(PostLatestCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PostLatestCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(PostLatestCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_postLatestCache_MultiGet(t *testing.T) {
	c := newPostLatestCache()
	defer c.Close()

	var testData []*model.PostLatest
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.PostLatest))
	}

	err := c.ICache.(PostLatestCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PostLatestCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.PostLatest))
	}
}

func Test_postLatestCache_MultiSet(t *testing.T) {
	c := newPostLatestCache()
	defer c.Close()

	var testData []*model.PostLatest
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.PostLatest))
	}

	err := c.ICache.(PostLatestCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_postLatestCache_Del(t *testing.T) {
	c := newPostLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostLatest)
	err := c.ICache.(PostLatestCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_postLatestCache_SetCacheWithNotFound(t *testing.T) {
	c := newPostLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PostLatest)
	err := c.ICache.(PostLatestCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewPostLatestCache(t *testing.T) {
	c := NewPostLatestCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
