package cache

import (
	"testing"
	"time"

	"creation/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserLikeCache() *gotest.Cache {
	record1 := &model.UserLike{}
	record1.ID = 1
	record2 := &model.UserLike{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserLikeCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userLikeCache_Set(t *testing.T) {
	c := newUserLikeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserLike)
	err := c.ICache.(UserLikeCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserLikeCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userLikeCache_Get(t *testing.T) {
	c := newUserLikeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserLike)
	err := c.ICache.(UserLikeCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserLikeCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserLikeCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userLikeCache_MultiGet(t *testing.T) {
	c := newUserLikeCache()
	defer c.Close()

	var testData []*model.UserLike
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserLike))
	}

	err := c.ICache.(UserLikeCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserLikeCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserLike))
	}
}

func Test_userLikeCache_MultiSet(t *testing.T) {
	c := newUserLikeCache()
	defer c.Close()

	var testData []*model.UserLike
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserLike))
	}

	err := c.ICache.(UserLikeCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userLikeCache_Del(t *testing.T) {
	c := newUserLikeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserLike)
	err := c.ICache.(UserLikeCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userLikeCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserLikeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserLike)
	err := c.ICache.(UserLikeCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserLikeCache(t *testing.T) {
	c := NewUserLikeCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
