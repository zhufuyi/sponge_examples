package cache

import (
	"testing"
	"time"

	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

type userTokenData struct {
	ID    uint64
	Key   interface{}
	Value interface{}
}

func newUserTokenCache() *gotest.Cache {
	// change the type of the value before testing
	var (
		key uint64 = 1
		val string = "bar1"
	)

	record1 := &userTokenData{ID: 1, Key: key, Value: val}
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserTokenCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userTokenCache_Set(t *testing.T) {
	c := newUserTokenCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*userTokenData)
	key := record.Key.(uint64)
	value := record.Value.(string)
	err := c.ICache.(UserTokenCache).Set(c.Ctx, key, value, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userTokenCache_Get(t *testing.T) {
	c := newUserTokenCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*userTokenData)
	key := record.Key.(uint64)
	value := record.Value.(string)
	err := c.ICache.(UserTokenCache).Set(c.Ctx, key, value, time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserTokenCache).Get(c.Ctx, key)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, value, got)
}

func Test_userTokenCache_Del(t *testing.T) {
	c := newUserTokenCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*userTokenData)
	key := record.Key.(uint64)
	err := c.ICache.(UserTokenCache).Del(c.Ctx, key)
	if err != nil {
		t.Fatal(err)
	}
}
