package model

import (
	"github.com/tj/assert"
	"testing"
)

func TestTodo_Table(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := Todo{}.Table()
		assert.Equal(t, "todos", result)
	})
}

func TestTodo_PrimaryKey(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := Todo{}.PrimaryKey()
		assert.Len(t, result.Columns(), 1)
		assert.Equal(t, "id", result.Columns()[0])
	})
}

func TestTodo_Indexes(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := Todo{}.Indexes()
		assert.Len(t, result, 1)

		idx := result[0]
		columns := idx.Columns()
		assert.Equal(t, "user_id_idx", idx.Name())
		assert.Len(t, columns, 1)
		assert.Equal(t, "user_id", columns[0])
	})
}
