package config

import (
	"testing"

	"coupon/configs"

	"github.com/zhufuyi/sponge/pkg/gofile"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	configFile := configs.Path("coupon.yml")
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

	// set nil
	Set(nil)
	defer func() {
		recover()
	}()
	Get()
}

func TestInitNacos(t *testing.T) {
	configFile := configs.Path("coupon_cc.yml")
	_, err := NewCenter(configFile)
	if gofile.IsExists(configFile) {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}
}
