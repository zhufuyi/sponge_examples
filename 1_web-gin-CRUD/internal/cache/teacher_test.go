package cache

import (
	"testing"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newTeacherCache() *gotest.Cache {
	record1 := &model.Teacher{}
	record1.ID = 1
	record2 := &model.Teacher{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewTeacherCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_teacherCache_Set(t *testing.T) {
	c := newTeacherCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teacher)
	err := c.ICache.(TeacherCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(TeacherCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_teacherCache_Get(t *testing.T) {
	c := newTeacherCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teacher)
	err := c.ICache.(TeacherCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(TeacherCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(TeacherCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_teacherCache_MultiGet(t *testing.T) {
	c := newTeacherCache()
	defer c.Close()

	var testData []*model.Teacher
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Teacher))
	}

	err := c.ICache.(TeacherCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(TeacherCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Teacher))
	}
}

func Test_teacherCache_MultiSet(t *testing.T) {
	c := newTeacherCache()
	defer c.Close()

	var testData []*model.Teacher
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Teacher))
	}

	err := c.ICache.(TeacherCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_teacherCache_Del(t *testing.T) {
	c := newTeacherCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teacher)
	err := c.ICache.(TeacherCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_teacherCache_SetCacheWithNotFound(t *testing.T) {
	c := newTeacherCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Teacher)
	err := c.ICache.(TeacherCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewTeacherCache(t *testing.T) {
	c := NewTeacherCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
