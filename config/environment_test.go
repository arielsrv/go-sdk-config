package config_test

import (
	"testing"

	"github.com/go-sdk-config/config"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	actual := config.IsEmpty("")
	assert.True(t, actual)
}

func TestIsNotEmpty(t *testing.T) {
	actual := config.IsEmpty("value")
	assert.False(t, actual)
}

func TestGetScope(t *testing.T) {
	t.Setenv("SCOPE", "test")
	actual := config.GetScope()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "test", actual)
}

func TestGetEnv(t *testing.T) {
	actual := config.GetEnv()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "local", actual)
}

func TestGetEnv_Custom(t *testing.T) {
	t.Setenv("app.env", "staging")
	actual := config.GetEnv()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "staging", actual)
}

func TestGetEnv_Custom_(t *testing.T) {
	t.Setenv("app_env", "staging")
	actual := config.GetEnv()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "staging", actual)
}

func TestGetEnv_Prod(t *testing.T) {
	t.Setenv("SCOPE", "prod")
	actual := config.GetEnv()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "prod", actual)
	assert.True(t, config.IsProd())
}

func TestIsDev(t *testing.T) {
	assert.True(t, config.IsLocal())
}
