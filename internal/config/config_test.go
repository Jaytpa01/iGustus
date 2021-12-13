package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	err := ReadConfig()

	assert.Nil(t, err)

	if _, ok := Config.Models["post"]; !ok {
		t.Error("post not found")
	}

	assert.Equal(t, float32(0.7), Config.Models["post"].Temperature, Config)
	assert.Equal(t, 38, Config.Models["post"].MaxTokens, Config)
}
