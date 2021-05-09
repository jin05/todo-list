package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tj/assert"
	"regexp"
	"testing"
	"time"
	"todo-list/app/infrastructure/database"
	test_database "todo-list/app/library/test/database"
)

func newUserRepository(conn *database.Connection) *userRepository {
	return NewUserRepository(conn).(*userRepository)
}

func TestNewUserRepository(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := NewUserRepository(nil).(*userRepository)
		assert.NotNil(t, result)
		assert.Nil(t, result.conn)
	})
}

func Test_userRepository_GetByAuthID(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			userRepo := newUserRepository(conn)

			userID := int64(1)
			userName := "user_name"
			authID := "auth_id"
			email := "email"
			now := time.Time{}
			mock.ExpectQuery(regexp.QuoteMeta("")).WithArgs(authID).
				WillReturnRows(sqlmock.NewRows([]string{
					"id", "user_name", "auth_id", "email", "created_at", "updated_at",
				}).AddRow(userID, userName, authID, email, now, now))

			result, err := userRepo.GetByAuthID(authID)
			assert.Nil(t, err)
			assert.Equal(t, userID, result.UserID)
			assert.Equal(t, userName, result.UserName)
			assert.Equal(t, authID, result.AuthID)
			assert.Equal(t, email, result.Email)
		})
		assert.Nil(t, err)
	})
}

func Test_userRepository_Save(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			userRepo := newUserRepository(conn)

			authID := "auth_id"
			name := "name"
			email := "email"
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(
				"INSERT INTO `users` (`user_name`,`auth_id`,`email`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
				WithArgs(name, authID, email, test_database.AnyTime{}, test_database.AnyTime{}).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			result, err := userRepo.Save(authID, name, email)
			assert.Nil(t, err)
			assert.Equal(t, int64(1), result.UserID)
			assert.Equal(t, authID, result.AuthID)
			assert.Equal(t, name, result.UserName)
			assert.Equal(t, email, result.Email)
		})
		assert.Nil(t, err)
	})
}
