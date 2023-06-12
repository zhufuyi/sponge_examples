package cache

import (
	"testing"
	"time"

	"community/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newRelationNumCache() *gotest.Cache {
	record1 := &model.RelationNum{}
	record1.ID = 1
	record2 := &model.RelationNum{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewRelationNumCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_relationNumCache_Set(t *testing.T) {
	c := newRelationNumCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.RelationNum)
	err := c.ICache.(RelationNumCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(RelationNumCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_relationNumCache_Get(t *testing.T) {
	c := newRelationNumCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.RelationNum)
	err := c.ICache.(RelationNumCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(RelationNumCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(RelationNumCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_relationNumCache_MultiGet(t *testing.T) {
	c := newRelationNumCache()
	defer c.Close()

	var testData []*model.RelationNum
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.RelationNum))
	}

	err := c.ICache.(RelationNumCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(RelationNumCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.RelationNum))
	}
}

func Test_relationNumCache_MultiSet(t *testing.T) {
	c := newRelationNumCache()
	defer c.Close()

	var testData []*model.RelationNum
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.RelationNum))
	}

	err := c.ICache.(RelationNumCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_relationNumCache_Del(t *testing.T) {
	c := newRelationNumCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.RelationNum)
	err := c.ICache.(RelationNumCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_relationNumCache_SetCacheWithNotFound(t *testing.T) {
	c := newRelationNumCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.RelationNum)
	err := c.ICache.(RelationNumCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewRelationNumCache(t *testing.T) {
	c := NewRelationNumCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
