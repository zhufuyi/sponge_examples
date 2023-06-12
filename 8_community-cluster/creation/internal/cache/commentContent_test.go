package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newCommentContentCache() *gotest.Cache {
	record1 := &model.CommentContent{}
	record1.ID = 1
	record2 := &model.CommentContent{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCommentContentCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_commentContentCache_Set(t *testing.T) {
	c := newCommentContentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentContent)
	err := c.ICache.(CommentContentCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CommentContentCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_commentContentCache_Get(t *testing.T) {
	c := newCommentContentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentContent)
	err := c.ICache.(CommentContentCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentContentCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CommentContentCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_commentContentCache_MultiGet(t *testing.T) {
	c := newCommentContentCache()
	defer c.Close()

	var testData []*model.CommentContent
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CommentContent))
	}

	err := c.ICache.(CommentContentCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentContentCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.CommentContent))
	}
}

func Test_commentContentCache_MultiSet(t *testing.T) {
	c := newCommentContentCache()
	defer c.Close()

	var testData []*model.CommentContent
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CommentContent))
	}

	err := c.ICache.(CommentContentCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentContentCache_Del(t *testing.T) {
	c := newCommentContentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentContent)
	err := c.ICache.(CommentContentCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentContentCache_SetCacheWithNotFound(t *testing.T) {
	c := newCommentContentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CommentContent)
	err := c.ICache.(CommentContentCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCommentContentCache(t *testing.T) {
	c := NewCommentContentCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
