package cache

import (
	"testing"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

type emailVerifyCodeData struct {
	ID    uint64
	Key   interface{}
	Value interface{}
}

func newEmailVerifyCodeCache() *gotest.Cache {
	// change the type of the value before testing
	var (
		key string = "foo1"
		val string = "bar1"
	)

	record1 := &emailVerifyCodeData{ID: 1, Key: key, Value: val}
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewEmailVerifyCodeCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_emailVerifyCodeCache_Set(t *testing.T) {
	c := newEmailVerifyCodeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*emailVerifyCodeData)
	key := record.Key.(string)
	value := record.Value.(string)
	err := c.ICache.(EmailVerifyCodeCache).Set(c.Ctx, key, value, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_emailVerifyCodeCache_Get(t *testing.T) {
	c := newEmailVerifyCodeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*emailVerifyCodeData)
	key := record.Key.(string)
	value := record.Value.(string)
	err := c.ICache.(EmailVerifyCodeCache).Set(c.Ctx, key, value, time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(EmailVerifyCodeCache).Get(c.Ctx, key)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, value, got)
}

func Test_emailVerifyCodeCache_Del(t *testing.T) {
	c := newEmailVerifyCodeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*emailVerifyCodeData)
	key := record.Key.(string)
	err := c.ICache.(EmailVerifyCodeCache).Del(c.Ctx, key)
	if err != nil {
		t.Fatal(err)
	}
}
