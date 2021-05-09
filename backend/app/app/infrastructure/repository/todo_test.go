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

func newTestTodoRepository(conn *database.Connection) *todoRepository {
	return NewTodoRepository(conn).(*todoRepository)
}

func TestNewTodoRepository(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := NewTodoRepository(nil).(*todoRepository)
		assert.NotNil(t, result)
		assert.Nil(t, result.conn)
	})
}

func Test_todoRepository_Get(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			todoID := int64(2)
			now := time.Time{}
			mock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `todos` "+
					"WHERE `todos`.`id` = ? "+
					"AND `todos`.`user_id` = ? "+
					"ORDER BY `todos`.`id` LIMIT 1")).
				WithArgs(todoID, userID).
				WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"user_id",
					"title",
					"content",
					"checked",
					"created_at",
					"updated_at",
				}).AddRow(
					todoID,
					userID,
					"title",
					"content",
					true,
					now,
					now,
				))

			result, err := todoRepo.Get(userID, todoID)
			assert.Nil(t, err)
			assert.Equal(t, todoID, result.TodoID)
			assert.Equal(t, userID, result.UserID)
			assert.Equal(t, "title", result.Title)
			assert.Equal(t, "content", result.Content)
			assert.True(t, result.Checked)
		})
		assert.Nil(t, err)
	})
	t.Run("error test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			todoID := int64(2)
			mock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `todos` "+
					"WHERE `todos`.`id` = ? "+
					"AND `todos`.`user_id` = ? "+
					"ORDER BY `todos`.`id` LIMIT 1")).
				WithArgs(todoID, userID).
				WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"user_id",
					"title",
					"content",
					"checked",
					"created_at",
					"updated_at",
				}))

			result, err := todoRepo.Get(userID, todoID)
			assert.Error(t, err)
			assert.Nil(t, result)
		})
		assert.Nil(t, err)
	})
}

func Test_todoRepository_List(t *testing.T) {
	t.Run("title query test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			keywords := []string{"Go", "Java", "Ruby"}

			now := time.Time{}
			mock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `todos` " +
					"WHERE user_id = ? " +
					"AND title LIKE '%Go%' " +
					"AND title LIKE '%Java%' " +
					"AND title LIKE '%Ruby%'")).
				WithArgs(userID).
				WillReturnRows(sqlmock.NewRows([]string{
					"id", "user_id", "title", "content", "checked", "created_at", "updated_at"}).
					AddRow(int64(1), userID, "title1", "content1", true, now, now).
					AddRow(int64(2), userID, "title2", "content2", false, now, now).
					AddRow(int64(3), userID, "title3", "content3", true, now, now))

			list, err := todoRepo.List(userID, keywords, "title")

			assert.Nil(t, err)
			assert.Len(t, list, 3)
			assert.Equal(t, int64(1), list[0].TodoID)
			assert.Equal(t, "title1", list[0].Title)
			assert.Equal(t, int64(2), list[1].TodoID)
			assert.Equal(t, "title2", list[1].Title)
			assert.Equal(t, int64(3), list[2].TodoID)
			assert.Equal(t, "title3", list[2].Title)
		})
		assert.Nil(t, err)
	})

	t.Run("content query test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			keywords := []string{"Go", "Java", "Ruby"}

			now := time.Time{}
			mock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `todos` " +
					"WHERE user_id = ? " +
					"AND content LIKE '%Go%' " +
					"AND content LIKE '%Java%' " +
					"AND content LIKE '%Ruby%'")).
				WithArgs(userID).
				WillReturnRows(sqlmock.NewRows([]string{
					"id", "user_id", "title", "content", "checked", "created_at", "updated_at"}).
					AddRow(int64(1), userID, "title1", "content1", true, now, now).
					AddRow(int64(2), userID, "title2", "content2", false, now, now).
					AddRow(int64(3), userID, "title3", "content3", true, now, now))

			list, err := todoRepo.List(userID, keywords, "content")

			assert.Nil(t, err)
			assert.Len(t, list, 3)
			assert.Equal(t, int64(1), list[0].TodoID)
			assert.Equal(t, "title1", list[0].Title)
			assert.Equal(t, int64(2), list[1].TodoID)
			assert.Equal(t, "title2", list[1].Title)
			assert.Equal(t, int64(3), list[2].TodoID)
			assert.Equal(t, "title3", list[2].Title)
		})
		assert.Nil(t, err)
	})

	t.Run("keywords empty test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			var keywords []string

			now := time.Time{}
			mock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `todos` WHERE user_id = ?")).
				WithArgs(userID).
				WillReturnRows(sqlmock.NewRows([]string{
					"id", "user_id", "title", "content", "checked", "created_at", "updated_at"}).
					AddRow(int64(1), userID, "title1", "content1", true, now, now).
					AddRow(int64(2), userID, "title2", "content2", false, now, now).
					AddRow(int64(3), userID, "title3", "content3", true, now, now))

			list, err := todoRepo.List(userID, keywords, "content")

			assert.Nil(t, err)
			assert.Len(t, list, 3)
			assert.Equal(t, int64(1), list[0].TodoID)
			assert.Equal(t, "title1", list[0].Title)
			assert.Equal(t, int64(2), list[1].TodoID)
			assert.Equal(t, "title2", list[1].Title)
			assert.Equal(t, int64(3), list[2].TodoID)
			assert.Equal(t, "title3", list[2].Title)
		})
		assert.Nil(t, err)
	})

	t.Run("searchTarget empty test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {
			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			keywords := []string{"Go", "Java", "Ruby"}

			now := time.Time{}
			mock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `todos` WHERE user_id = ?")).
				WithArgs(userID).
				WillReturnRows(sqlmock.NewRows([]string{
					"id", "user_id", "title", "content", "checked", "created_at", "updated_at"}).
					AddRow(int64(1), userID, "title1", "content1", true, now, now).
					AddRow(int64(2), userID, "title2", "content2", false, now, now).
					AddRow(int64(3), userID, "title3", "content3", true, now, now))

			list, err := todoRepo.List(userID, keywords, "")

			assert.Nil(t, err)
			assert.Len(t, list, 3)
			assert.Equal(t, int64(1), list[0].TodoID)
			assert.Equal(t, "title1", list[0].Title)
			assert.Equal(t, int64(2), list[1].TodoID)
			assert.Equal(t, "title2", list[1].Title)
			assert.Equal(t, int64(3), list[2].TodoID)
			assert.Equal(t, "title3", list[2].Title)
		})
		assert.Nil(t, err)
	})
}

func Test_todoRepository_Save(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {

			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			title := "title"
			content := "content"

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(
				"INSERT INTO `todos` (`user_id`,`title`,`content`,`checked`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
				WithArgs(userID, title, content, false, test_database.AnyTime{}, test_database.AnyTime{}).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			resutl, err := todoRepo.Save(userID, title, content)
			assert.Nil(t, err)
			assert.Equal(t, int64(1), resutl.TodoID)
			assert.Equal(t, userID, resutl.UserID)
			assert.Equal(t, title, resutl.Title)
			assert.Equal(t, content, resutl.Content)
			assert.False(t, resutl.Checked)
		})
		assert.Nil(t, err)
	})
}

func Test_todoRepository_Update(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {

			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			todoID := int64(2)
			title := "title"
			content := "content"

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(
				"UPDATE `todos` SET `title`=?,`content`=?,`checked`=?,`updated_at`=? "+
					"WHERE `todos`.`id` = ? AND `todos`.`user_id` = ?")).
				WithArgs(title, content, true, test_database.AnyTime{}, todoID, userID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			err := todoRepo.Update(userID, todoID, title, content, true)
			assert.Nil(t, err)
		})
		assert.Nil(t, err)
	})
}

func Test_todoRepository_Delete(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {

			todoRepo := newTestTodoRepository(conn)

			userID := int64(1)
			todoID := int64(2)

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(
				"DELETE FROM `todos` WHERE `todos`.`id` = ? AND `todos`.`user_id` = ?")).
				WithArgs(todoID, userID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			err := todoRepo.Delete(userID, todoID)
			assert.Nil(t, err)
		})
		assert.Nil(t, err)
	})
}
