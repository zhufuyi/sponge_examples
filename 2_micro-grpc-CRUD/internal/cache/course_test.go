package cache

import (
	"testing"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newCourseCache() *gotest.Cache {
	record1 := &model.Course{}
	record1.ID = 1
	record2 := &model.Course{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCourseCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_courseCache_Set(t *testing.T) {
	c := newCourseCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Course)
	err := c.ICache.(CourseCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CourseCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_courseCache_Get(t *testing.T) {
	c := newCourseCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Course)
	err := c.ICache.(CourseCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CourseCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CourseCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_courseCache_MultiGet(t *testing.T) {
	c := newCourseCache()
	defer c.Close()

	var testData []*model.Course
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Course))
	}

	err := c.ICache.(CourseCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CourseCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Course))
	}
}

func Test_courseCache_MultiSet(t *testing.T) {
	c := newCourseCache()
	defer c.Close()

	var testData []*model.Course
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Course))
	}

	err := c.ICache.(CourseCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_courseCache_Del(t *testing.T) {
	c := newCourseCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Course)
	err := c.ICache.(CourseCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_courseCache_SetCacheWithNotFound(t *testing.T) {
	c := newCourseCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Course)
	err := c.ICache.(CourseCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCourseCache(t *testing.T) {
	c := NewCourseCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
