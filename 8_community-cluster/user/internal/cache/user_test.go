package cache

import (
	"testing"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserCache() *gotest.Cache {
	record1 := &model.User{}
	record1.ID = 1
	record2 := &model.User{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userCache_Set(t *testing.T) {
	c := newUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.User)
	err := c.ICache.(UserCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userCache_Get(t *testing.T) {
	c := newUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.User)
	err := c.ICache.(UserCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userCache_MultiGet(t *testing.T) {
	c := newUserCache()
	defer c.Close()

	var testData []*model.User
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.User))
	}

	err := c.ICache.(UserCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.User))
	}
}

func Test_userCache_MultiSet(t *testing.T) {
	c := newUserCache()
	defer c.Close()

	var testData []*model.User
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.User))
	}

	err := c.ICache.(UserCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userCache_Del(t *testing.T) {
	c := newUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.User)
	err := c.ICache.(UserCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.User)
	err := c.ICache.(UserCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserCache(t *testing.T) {
	c := NewUserCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
