package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newCommentCache() *gotest.Cache {
	record1 := &model.Comment{}
	record1.ID = 1
	record2 := &model.Comment{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCommentCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_commentCache_Set(t *testing.T) {
	c := newCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Comment)
	err := c.ICache.(CommentCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CommentCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_commentCache_Get(t *testing.T) {
	c := newCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Comment)
	err := c.ICache.(CommentCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CommentCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_commentCache_MultiGet(t *testing.T) {
	c := newCommentCache()
	defer c.Close()

	var testData []*model.Comment
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Comment))
	}

	err := c.ICache.(CommentCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CommentCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Comment))
	}
}

func Test_commentCache_MultiSet(t *testing.T) {
	c := newCommentCache()
	defer c.Close()

	var testData []*model.Comment
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Comment))
	}

	err := c.ICache.(CommentCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentCache_Del(t *testing.T) {
	c := newCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Comment)
	err := c.ICache.(CommentCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_commentCache_SetCacheWithNotFound(t *testing.T) {
	c := newCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Comment)
	err := c.ICache.(CommentCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCommentCache(t *testing.T) {
	c := NewCommentCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
