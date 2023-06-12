package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newCommentLatestCache() *gotest.Cache {
	record1 := &model.CommentLatest{}
	record1.ID = 1
	record2 := &model.CommentLatest{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCommentLatestCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_commentLatestCache_Set(t *testing.T) {
	c := newCommentLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentLatest)
	err := c.ICache.(CommentLatestCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CommentLatestCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_commentLatestCache_Get(t *testing.T) {
	c := newCommentLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentLatest)
	err := c.ICache.(CommentLatestCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentLatestCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CommentLatestCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_commentLatestCache_MultiGet(t *testing.T) {
	c := newCommentLatestCache()
	defer c.Close()

	var testData []*model.CommentLatest
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CommentLatest))
	}

	err := c.ICache.(CommentLatestCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentLatestCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.CommentLatest))
	}
}

func Test_commentLatestCache_MultiSet(t *testing.T) {
	c := newCommentLatestCache()
	defer c.Close()

	var testData []*model.CommentLatest
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CommentLatest))
	}

	err := c.ICache.(CommentLatestCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentLatestCache_Del(t *testing.T) {
	c := newCommentLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentLatest)
	err := c.ICache.(CommentLatestCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentLatestCache_SetCacheWithNotFound(t *testing.T) {
	c := newCommentLatestCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentLatest)
	err := c.ICache.(CommentLatestCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCommentLatestCache(t *testing.T) {
	c := NewCommentLatestCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
