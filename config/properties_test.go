package config_test

import (
	"testing"

	"github.com/go-sdk-config/config"
	"github.com/stretchr/testify/assert"
)

func TestAppConfig(t *testing.T) {
	actual := config.NewBuilder().
		WithFile("config.yml").
		WithFolder("../resources").
		Build()

	assert.NoError(t, actual.Err)

	stringValue := config.String("app_name")
	assert.Equal(t, "go-sdk-config", stringValue)

	stringValue = config.String("missing")
	assert.Equal(t, "", stringValue)

	boolValue := config.TryBool("enable", true)
	assert.True(t, boolValue)

	booleanKey := config.TryBool("boolean.key", false)
	assert.True(t, booleanKey)

	boolValue = config.TryBool("logger", true)
	assert.True(t, boolValue)

	intValue := config.TryInt("missing threads", 1)
	assert.Equal(t, 1, intValue)

	intValue = config.TryInt("threads", 1)
	assert.Equal(t, 1, intValue)

	intKey := config.TryInt("int.key", 1)
	assert.Equal(t, 10, intKey)
}
