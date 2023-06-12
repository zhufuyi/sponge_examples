package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserCommentCache() *gotest.Cache {
	record1 := &model.UserComment{}
	record1.ID = 1
	record2 := &model.UserComment{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserCommentCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userCommentCache_Set(t *testing.T) {
	c := newUserCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserComment)
	err := c.ICache.(UserCommentCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserCommentCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userCommentCache_Get(t *testing.T) {
	c := newUserCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserComment)
	err := c.ICache.(UserCommentCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserCommentCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserCommentCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userCommentCache_MultiGet(t *testing.T) {
	c := newUserCommentCache()
	defer c.Close()

	var testData []*model.UserComment
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserComment))
	}

	err := c.ICache.(UserCommentCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserCommentCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserComment))
	}
}

func Test_userCommentCache_MultiSet(t *testing.T) {
	c := newUserCommentCache()
	defer c.Close()

	var testData []*model.UserComment
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserComment))
	}

	err := c.ICache.(UserCommentCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userCommentCache_Del(t *testing.T) {
	c := newUserCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserComment)
	err := c.ICache.(UserCommentCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userCommentCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserCommentCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserComment)
	err := c.ICache.(UserCommentCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserCommentCache(t *testing.T) {
	c := NewUserCommentCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
