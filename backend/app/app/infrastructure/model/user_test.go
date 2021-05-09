package model

import (
	"github.com/tj/assert"
	"testing"
)

func TestUser_Table(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := User{}.Table()
		assert.Equal(t, "users", result)
	})
}

func TestUser_PrimaryKey(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := User{}.PrimaryKey()
		assert.Len(t, result.Columns(), 1)
		assert.Equal(t, "id", result.Columns()[0])
	})
}

func TestUser_Indexes(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := User{}.Indexes()
		assert.Len(t, result, 2)

		idx := result[0]
		columns := idx.Columns()
		assert.Equal(t, "auth_id_idx", idx.Name())
		assert.Len(t, columns, 1)
		assert.Equal(t, "auth_id", columns[0])

		idx = result[1]
		columns = idx.Columns()
		assert.Equal(t, "email_idx", idx.Name())
		assert.Len(t, columns, 1)
		assert.Equal(t, "email", columns[0])
	})
}
