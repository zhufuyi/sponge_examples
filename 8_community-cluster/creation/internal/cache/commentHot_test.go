package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newCommentHotCache() *gotest.Cache {
	record1 := &model.CommentHot{}
	record1.ID = 1
	record2 := &model.CommentHot{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCommentHotCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_commentHotCache_Set(t *testing.T) {
	c := newCommentHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentHot)
	err := c.ICache.(CommentHotCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CommentHotCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_commentHotCache_Get(t *testing.T) {
	c := newCommentHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentHot)
	err := c.ICache.(CommentHotCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentHotCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CommentHotCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_commentHotCache_MultiGet(t *testing.T) {
	c := newCommentHotCache()
	defer c.Close()

	var testData []*model.CommentHot
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CommentHot))
	}

	err := c.ICache.(CommentHotCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentHotCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.CommentHot))
	}
}

func Test_commentHotCache_MultiSet(t *testing.T) {
	c := newCommentHotCache()
	defer c.Close()

	var testData []*model.CommentHot
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CommentHot))
	}

	err := c.ICache.(CommentHotCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentHotCache_Del(t *testing.T) {
	c := newCommentHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentHot)
	err := c.ICache.(CommentHotCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentHotCache_SetCacheWithNotFound(t *testing.T) {
	c := newCommentHotCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentHot)
	err := c.ICache.(CommentHotCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCommentHotCache(t *testing.T) {
	c := NewCommentHotCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
