package database

import (
	"github.com/tj/assert"
	"testing"
	"todo-list/app/config"
)

func TestNewDB(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		conf := &config.Config{
			DB: config.DBConfig{
				User:   "user",
				Pass:   "pass",
				Host:   "host",
				Port:   3306,
				DBName: "dbname",
			},
		}

		db, err := NewDB(conf)
		assert.Error(t, err)
		assert.Nil(t, db)
	})
}
