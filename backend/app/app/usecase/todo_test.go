package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
	"testing"
	"todo-list/app/domain"
	mock_repository "todo-list/app/library/test/mock/infrastructure/repository"
)

func newTodoUseCase(ctrl *gomock.Controller) *todoUseCase {
	todoRepository := mock_repository.NewMockTodoRepository(ctrl)
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	return NewTodoUseCase(todoRepository, userRepository).(*todoUseCase)
}

func newUser() *domain.User {
	return &domain.User{
		UserID:   1,
		UserName: "user_name",
		AuthID:   "auth_id",
		Email:    "email",
	}
}

func Test_todoUseCase_Get(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		user := newUser()
		todo := &domain.Todo{
			TodoID:  2,
			UserID:  1,
			Title:   "title",
			Content: "content",
			Checked: false,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := newTodoUseCase(ctrl)

		userRepository := useCase.userRepository.(*mock_repository.MockUserRepository)
		todoRepository := useCase.todoRepository.(*mock_repository.MockTodoRepository)

		gomock.InOrder(
			userRepository.EXPECT().GetByAuthID(user.AuthID).Return(user, nil),
			todoRepository.EXPECT().Get(user.UserID, todo.TodoID).Return(todo, nil),
		)

		result, err := useCase.Get(user.AuthID, todo.TodoID)
		assert.Equal(t, todo.TodoID, result.TodoID)
		assert.Equal(t, todo.UserID, result.UserID)
		assert.Equal(t, todo.Title, result.Title)
		assert.Equal(t, todo.Content, result.Content)
		assert.Equal(t, todo.Checked, result.Checked)
		assert.Nil(t, err)
	})
}

func Test_todoUseCase_Create(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		user := newUser()
		todo := &domain.Todo{
			TodoID:  2,
			UserID:  1,
			Title:   "title",
			Content: "content",
			Checked: false,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := newTodoUseCase(ctrl)

		userRepository := useCase.userRepository.(*mock_repository.MockUserRepository)
		todoRepository := useCase.todoRepository.(*mock_repository.MockTodoRepository)

		gomock.InOrder(
			userRepository.EXPECT().GetByAuthID(user.AuthID).Return(user, nil),
			todoRepository.EXPECT().Save(user.UserID, todo.Title, todo.Content).Return(todo, nil),
		)

		result, err := useCase.Create(user.AuthID, todo.Title, todo.Content)
		assert.Equal(t, todo.TodoID, result.TodoID)
		assert.Equal(t, todo.UserID, result.UserID)
		assert.Equal(t, todo.Title, result.Title)
		assert.Equal(t, todo.Content, result.Content)
		assert.Equal(t, todo.Checked, result.Checked)
		assert.Nil(t, err)
	})
}

func Test_todoUseCase_Update(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		user := newUser()
		todo := &domain.Todo{
			TodoID:  2,
			UserID:  1,
			Title:   "title",
			Content: "content",
			Checked: false,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := newTodoUseCase(ctrl)

		userRepository := useCase.userRepository.(*mock_repository.MockUserRepository)
		todoRepository := useCase.todoRepository.(*mock_repository.MockTodoRepository)

		title := "title_update"
		content := "content_update"
		checked := true

		gomock.InOrder(
			userRepository.EXPECT().GetByAuthID(user.AuthID).Return(user, nil),
			todoRepository.EXPECT().Get(user.UserID, todo.TodoID).Return(todo, nil),
			todoRepository.EXPECT().Update(user.UserID, todo.TodoID, title, content, checked).Return(nil),
		)

		result, err := useCase.Update(user.AuthID, todo.TodoID, title, content, checked)
		assert.Equal(t, todo.TodoID, result.TodoID)
		assert.Equal(t, todo.UserID, result.UserID)
		assert.Equal(t, title, result.Title)
		assert.Equal(t, content, result.Content)
		assert.Equal(t, checked, result.Checked)
		assert.Nil(t, err)
	})
}

func Test_todoUseCase_Delete(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		user := newUser()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := newTodoUseCase(ctrl)

		userRepository := useCase.userRepository.(*mock_repository.MockUserRepository)
		todoRepository := useCase.todoRepository.(*mock_repository.MockTodoRepository)

		todoID := int64(2)
		gomock.InOrder(
			userRepository.EXPECT().GetByAuthID(user.AuthID).Return(user, nil),
			todoRepository.EXPECT().Delete(user.UserID, todoID).Return(nil),
		)

		err := useCase.Delete(user.AuthID, todoID)
		assert.Nil(t, err)
	})
}

func Test_todoUseCase_List(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		user := newUser()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := newTodoUseCase(ctrl)

		userRepository := useCase.userRepository.(*mock_repository.MockUserRepository)
		todoRepository := useCase.todoRepository.(*mock_repository.MockTodoRepository)

		searchTarget := "title"
		keyword := "key word"
		keywords := []string{
			"key",
			"word",
		}
		list := []*domain.Todo{
			{
				TodoID:  1,
				UserID:  1,
				Title:   "title1",
				Content: "content1",
				Checked: false,
			},
			{
				TodoID:  2,
				UserID:  1,
				Title:   "title2",
				Content: "content2",
				Checked: false,
			},
			{
				TodoID:  3,
				UserID:  1,
				Title:   "title3",
				Content: "content3",
				Checked: false,
			},
		}
		gomock.InOrder(
			userRepository.EXPECT().GetByAuthID(user.AuthID).Return(user, nil),
			todoRepository.EXPECT().List(user.UserID, keywords, searchTarget).Return(list, nil),
		)

		result, err := useCase.List(user.AuthID, keyword, searchTarget)
		assert.Nil(t, err)
		for i, todo := range list {
			assert.Equal(t, todo.TodoID, result[i].TodoID)
			assert.Equal(t, todo.UserID, result[i].UserID)
			assert.Equal(t, todo.Title, result[i].Title)
			assert.Equal(t, todo.Content, result[i].Content)
			assert.Equal(t, todo.Checked, result[i].Checked)
		}
	})
}
