package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gofile"

	"eshop/order_gw/configs"
)

func TestInit(t *testing.T) {
	configFile := configs.Path("order_gw.yml")
	err := Init(configFile)
	if gofile.IsExists(configFile) {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}

	c := Get()
	assert.NotNil(t, c)

	str := Show()
	assert.NotEmpty(t, str)
	t.Log(str)

	// set nil
	Set(nil)
	defer func() {
		recover()
	}()
	Get()
}

func TestInitNacos(t *testing.T) {
	configFile := configs.Path("order_gw_cc.yml")
	_, err := NewCenter(configFile)
	if gofile.IsExists(configFile) {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}
}