package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newPostCache() *gotest.Cache {
	record1 := &model.Post{}
	record1.ID = 1
	record2 := &model.Post{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewPostCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_postCache_Set(t *testing.T) {
	c := newPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Post)
	err := c.ICache.(PostCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(PostCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_postCache_Get(t *testing.T) {
	c := newPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Post)
	err := c.ICache.(PostCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PostCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(PostCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_postCache_MultiGet(t *testing.T) {
	c := newPostCache()
	defer c.Close()

	var testData []*model.Post
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Post))
	}

	err := c.ICache.(PostCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PostCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Post))
	}
}

func Test_postCache_MultiSet(t *testing.T) {
	c := newPostCache()
	defer c.Close()

	var testData []*model.Post
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Post))
	}

	err := c.ICache.(PostCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_postCache_Del(t *testing.T) {
	c := newPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Post)
	err := c.ICache.(PostCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_postCache_SetCacheWithNotFound(t *testing.T) {
	c := newPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Post)
	err := c.ICache.(PostCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewPostCache(t *testing.T) {
	c := NewPostCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
