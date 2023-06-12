package cache

import (
	"testing"
	"time"

	"community/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserFollowerCache() *gotest.Cache {
	record1 := &model.UserFollower{}
	record1.ID = 1
	record2 := &model.UserFollower{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserFollowerCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userFollowerCache_Set(t *testing.T) {
	c := newUserFollowerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollower)
	err := c.ICache.(UserFollowerCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserFollowerCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userFollowerCache_Get(t *testing.T) {
	c := newUserFollowerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollower)
	err := c.ICache.(UserFollowerCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserFollowerCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserFollowerCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userFollowerCache_MultiGet(t *testing.T) {
	c := newUserFollowerCache()
	defer c.Close()

	var testData []*model.UserFollower
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserFollower))
	}

	err := c.ICache.(UserFollowerCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserFollowerCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserFollower))
	}
}

func Test_userFollowerCache_MultiSet(t *testing.T) {
	c := newUserFollowerCache()
	defer c.Close()

	var testData []*model.UserFollower
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserFollower))
	}

	err := c.ICache.(UserFollowerCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userFollowerCache_Del(t *testing.T) {
	c := newUserFollowerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollower)
	err := c.ICache.(UserFollowerCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userFollowerCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserFollowerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollower)
	err := c.ICache.(UserFollowerCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserFollowerCache(t *testing.T) {
	c := NewUserFollowerCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
