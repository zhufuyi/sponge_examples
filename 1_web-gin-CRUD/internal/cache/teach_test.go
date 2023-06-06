package cache

import (
	"testing"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newTeachCache() *gotest.Cache {
	record1 := &model.Teach{}
	record1.ID = 1
	record2 := &model.Teach{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewTeachCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_teachCache_Set(t *testing.T) {
	c := newTeachCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teach)
	err := c.ICache.(TeachCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(TeachCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_teachCache_Get(t *testing.T) {
	c := newTeachCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teach)
	err := c.ICache.(TeachCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(TeachCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(TeachCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_teachCache_MultiGet(t *testing.T) {
	c := newTeachCache()
	defer c.Close()

	var testData []*model.Teach
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Teach))
	}

	err := c.ICache.(TeachCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(TeachCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Teach))
	}
}

func Test_teachCache_MultiSet(t *testing.T) {
	c := newTeachCache()
	defer c.Close()

	var testData []*model.Teach
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Teach))
	}

	err := c.ICache.(TeachCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_teachCache_Del(t *testing.T) {
	c := newTeachCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teach)
	err := c.ICache.(TeachCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_teachCache_SetCacheWithNotFound(t *testing.T) {
	c := newTeachCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teach)
	err := c.ICache.(TeachCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewTeachCache(t *testing.T) {
	c := NewTeachCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
