package config

import (
	"github.com/tj/assert"
	"testing"
	"todo-list/app/library/test"
)

func TestNewConfig(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		envs := test.MakeTestEnvs()
		test.SetupEnvs(envs)
		defer test.ClearEnvs(envs)

		result, err := NewConfig()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.False(t, result.Server.IsProduction)
	})
}
