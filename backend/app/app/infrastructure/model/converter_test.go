package model

import (
	"github.com/tj/assert"
	"testing"
	"time"
)

func TestToTodoDomain(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		user := &User{
			ID:       1,
			UserName: "user_name",
			AuthID:   "auth_id",
			Email:    "email",
		}
		result := ToUserDomain(user)
		assert.Equal(t, user.ID, result.UserID)
		assert.Equal(t, user.UserName, result.UserName)
		assert.Equal(t, user.AuthID, result.AuthID)
		assert.Equal(t, user.Email, result.Email)
	})
}

func TestToTodoDomainList(t *testing.T) {
	list := []*Todo{
		{
			ID:        1,
			UserID:    2,
			Title:     "title",
			Content:   "content",
			Checked:   true,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			ID:        2,
			UserID:    3,
			Title:     "TTTTTT",
			Content:   "CCCCC",
			Checked:   false,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	result := ToTodoDomainList(list)
	assert.Len(t, result, 2)

	for i, todo := range list {
		assert.Equal(t, todo.ID, result[i].TodoID)
	}
}

func TestToUserDomain(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		todo := &Todo{
			ID:      1,
			UserID:  2,
			Title:   "title",
			Content: "content",
			Checked: true,
		}
		result := ToTodoDomain(todo)
		assert.Equal(t, todo.ID, result.TodoID)
		assert.Equal(t, todo.UserID, result.UserID)
		assert.Equal(t, todo.Title, result.Title)
		assert.Equal(t, todo.Content, result.Content)
		assert.Equal(t, todo.Checked, result.Checked)
	})
}
