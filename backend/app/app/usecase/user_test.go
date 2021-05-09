package usecase

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
	"testing"
	"todo-list/app/infrastructure/database"
	test_database "todo-list/app/library/test/database"
	mock_repository "todo-list/app/library/test/mock/infrastructure/repository"
)

func newUserUseCase(ctrl *gomock.Controller, conn *database.Connection) *userUseCase {
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	return NewUserUseCase(conn, userRepository).(*userUseCase)
}

func Test_userUseCase_GetByAuthID(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := newUserUseCase(ctrl, nil)

		userRepository := useCase.userRepository.(*mock_repository.MockUserRepository)

		user := newUser()
		gomock.InOrder(
			userRepository.EXPECT().GetByAuthID(user.AuthID).Return(user, nil),
		)

		result, err := useCase.GetByAuthID(user.AuthID)
		assert.Nil(t, err)
		assert.Equal(t, user.UserID, result.UserID)
		assert.Equal(t, user.AuthID, result.AuthID)
		assert.Equal(t, user.UserName, result.UserName)
		assert.Equal(t, user.Email, result.Email)
	})
}

func Test_userUseCase_Save(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		err := test_database.GetDBMock(func(conn *database.Connection, mock sqlmock.Sqlmock) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			useCase := newUserUseCase(ctrl, conn)

			userRepository := useCase.userRepository.(*mock_repository.MockUserRepository)
			user := newUser()

			mock.ExpectBegin()
			gomock.InOrder(
				userRepository.EXPECT().GetByAuthID(user.AuthID).Return(nil, nil),
				userRepository.EXPECT().Save(user.AuthID, user.UserName, user.Email).Return(user, nil),
			)
			mock.ExpectCommit()

			result, err := useCase.Save(user.AuthID, user.UserName, user.Email)
			assert.Nil(t, err)
			assert.Equal(t, user.UserID, result.UserID)
			assert.Equal(t, user.AuthID, result.AuthID)
			assert.Equal(t, user.UserName, result.UserName)
			assert.Equal(t, user.Email, result.Email)
		})
		assert.Nil(t, err)
	})
}
