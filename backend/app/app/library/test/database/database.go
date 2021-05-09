package test_database

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"time"
	"todo-list/app/infrastructure/database"
)

func GetDBMock(f func(conn *database.Connection, mock sqlmock.Sqlmock)) error {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return err
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT VERSION()`)).WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("5.7"))
	gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return err
	}

	conn := database.Connection{
		DB: gdb,
	}

	f(&conn, mock)
	return nil
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type Includes struct {
	Expects []interface{}
}

// Match satisfies sqlmock.Argument interface
func (i Includes) Match(v driver.Value) bool {
	for _, expect := range i.Expects {
		if reflect.DeepEqual(expect, v) {
			return true
		}
	}
	return false
}
