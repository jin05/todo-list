package api

import (
	"github.com/tj/assert"
	"testing"
)

func TestNewAPI(t *testing.T) {
	userApi := &userAPI{}
	todoAPI := &todoAPI{}

	result := NewAPI(userApi, todoAPI)
	assert.Equal(t, userApi, result.UserApi)
	assert.Equal(t, todoAPI, result.TodoAPI)
}
