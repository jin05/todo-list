package model

import (
	"github.com/tj/assert"
	"testing"
)

func TestGetModels(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		models := GetModels()

		assert.Len(t, models, 2)
		assert.Equal(t, &User{}, models[0])
		assert.Equal(t, &Todo{}, models[1])
	})
}
