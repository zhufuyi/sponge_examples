package cache

import (
	"testing"
	"time"

	"community/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserPostCache() *gotest.Cache {
	record1 := &model.UserPost{}
	record1.ID = 1
	record2 := &model.UserPost{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserPostCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userPostCache_Set(t *testing.T) {
	c := newUserPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserPost)
	err := c.ICache.(UserPostCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserPostCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userPostCache_Get(t *testing.T) {
	c := newUserPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserPost)
	err := c.ICache.(UserPostCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserPostCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserPostCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userPostCache_MultiGet(t *testing.T) {
	c := newUserPostCache()
	defer c.Close()

	var testData []*model.UserPost
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserPost))
	}

	err := c.ICache.(UserPostCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserPostCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserPost))
	}
}

func Test_userPostCache_MultiSet(t *testing.T) {
	c := newUserPostCache()
	defer c.Close()

	var testData []*model.UserPost
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserPost))
	}

	err := c.ICache.(UserPostCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userPostCache_Del(t *testing.T) {
	c := newUserPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserPost)
	err := c.ICache.(UserPostCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userPostCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserPostCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserPost)
	err := c.ICache.(UserPostCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserPostCache(t *testing.T) {
	c := NewUserPostCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
