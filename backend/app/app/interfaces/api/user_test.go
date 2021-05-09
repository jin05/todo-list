package api

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/app/domain"
	"todo-list/app/interfaces/middleware"
	mock_usecase "todo-list/app/library/test/mock/usecase"
)

func newUserAPI(ctrl *gomock.Controller) *userAPI {
	userUseCase := mock_usecase.NewMockUserUseCase(ctrl)
	return NewUserAPI(userUseCase).(*userAPI)
}

func newUser() *middleware.User {
	return &middleware.User{
		AuthID: "auth_id",
		Name:   "name",
		Email:  "email",
	}
}

func addUser(r *http.Request) *http.Request {
	ctx := r.Context()
	user := newUser()
	ctx = middleware.SetUser(ctx, user)
	return r.WithContext(ctx)
}

func TestNewUserAPI(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := newUserAPI(gomock.NewController(t))
		assert.NotNil(t, result)
	})
}

func Test_userAPI_Handler(t *testing.T) {
	t.Run("handler test", func(t *testing.T) {
		api := newUserAPI(gomock.NewController(t))

		r := httptest.NewRequest(
			http.MethodOptions,
			"https://test.com/user",
			bytes.NewBufferString(""),
		)
		w := httptest.NewRecorder()

		api.Handler(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("signup test", func(t *testing.T) {
		api := newUserAPI(gomock.NewController(t))
		userUseCase := api.userUseCase.(*mock_usecase.MockUserUseCase)
		userUseCase.EXPECT().Save(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&domain.User{}, nil)

		r := httptest.NewRequest(
			http.MethodPost,
			"https://test.com/user",
			bytes.NewBufferString(""),
		)
		r = addUser(r)
		w := httptest.NewRecorder()

		api.Handler(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func Test_userAPI_Signup(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		user := newUser()

		api := newUserAPI(gomock.NewController(t))
		userUseCase := api.userUseCase.(*mock_usecase.MockUserUseCase)
		userUseCase.EXPECT().Save(user.AuthID, user.Name, user.Email).
			Return(&domain.User{}, nil)

		r := httptest.NewRequest(
			http.MethodPost,
			"https://test.com/user",
			bytes.NewBufferString(""),
		)
		r = addUser(r)
		w := httptest.NewRecorder()

		api.Signup(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error test", func(t *testing.T) {
		user := newUser()

		api := newUserAPI(gomock.NewController(t))
		userUseCase := api.userUseCase.(*mock_usecase.MockUserUseCase)
		userUseCase.EXPECT().Save(user.AuthID, user.Name, user.Email).
			Return(nil, errors.New("test error"))

		r := httptest.NewRequest(
			http.MethodPost,
			"https://test.com/user",
			bytes.NewBufferString(""),
		)
		r = addUser(r)
		w := httptest.NewRecorder()

		api.Signup(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
