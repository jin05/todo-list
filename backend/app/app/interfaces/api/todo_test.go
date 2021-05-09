package api

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/tj/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/app/domain"
	mock_usecase "todo-list/app/library/test/mock/usecase"
)

func newTodoAPI(ctrl *gomock.Controller) *todoAPI {
	todoUseCase := mock_usecase.NewMockTodoUseCase(ctrl)
	return NewTodoAPI(todoUseCase).(*todoAPI)
}

func TestNewTodoAPI(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		result := newTodoAPI(gomock.NewController(t))
		assert.NotNil(t, result)
	})
}

func Test_todoAPI_Handler(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		api := newTodoAPI(gomock.NewController(t))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodOptions,
			"/todo",
			nil,
		)

		api.Handler(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func Test_todoAPI_Get(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		api := newTodoAPI(gomock.NewController(t))

		todo := &domain.Todo{
			TodoID:  1,
			UserID:  2,
			Title:   "title",
			Content: "content",
			Checked: false,
		}
		todoID := int64(1)
		todoUseCase := api.todoUseCase.(*mock_usecase.MockTodoUseCase)
		todoUseCase.EXPECT().Get("auth_id", todoID).Return(todo, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			"/todo",
			nil,
		)
		r = addUser(r)
		r = mux.SetURLVars(r, map[string]string{"todoID": "1"})

		api.Get(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "{\"TodoID\":1,\"UserID\":2,\"Title\":\"title\",\"Content\":\"content\",\"Checked\":false}\n", w.Body.String())
	})
}

func Test_todoAPI_Create(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		api := newTodoAPI(gomock.NewController(t))

		input := createInput{
			Title:   "title",
			Content: "content",
		}
		todoUseCase := api.todoUseCase.(*mock_usecase.MockTodoUseCase)
		todoUseCase.EXPECT().Create("auth_id", input.Title, input.Content).
			Return(&domain.Todo{
				TodoID:  1,
				UserID:  2,
				Title:   "title",
				Content: "content",
				Checked: false,
			}, nil)

		rBody, err := json.Marshal(input)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodPost,
			"/todo",
			bytes.NewBuffer(rBody),
		)
		r = addUser(r)

		api.Create(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "{\"TodoID\":1,\"UserID\":2,\"Title\":\"title\",\"Content\":\"content\",\"Checked\":false}\n", w.Body.String())
	})
}

func Test_todoAPI_Update(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		api := newTodoAPI(gomock.NewController(t))

		input := updateInput{
			TodoID:  1,
			Title:   "title",
			Content: "content",
			Checked: false,
		}
		todoUseCase := api.todoUseCase.(*mock_usecase.MockTodoUseCase)
		todoUseCase.EXPECT().Update("auth_id", input.TodoID, input.Title, input.Content, input.Checked).
			Return(&domain.Todo{
				TodoID:  1,
				UserID:  2,
				Title:   "title",
				Content: "content",
				Checked: false,
			}, nil)

		rBody, err := json.Marshal(input)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodPut,
			"/todo",
			bytes.NewBuffer(rBody),
		)
		r = addUser(r)

		api.Update(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "{\"TodoID\":1,\"UserID\":2,\"Title\":\"title\",\"Content\":\"content\",\"Checked\":false}\n", w.Body.String())
	})
}

func Test_todoAPI_Delete(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		api := newTodoAPI(gomock.NewController(t))

		input := deleteInput{TodoID: int64(1)}
		todoUseCase := api.todoUseCase.(*mock_usecase.MockTodoUseCase)
		todoUseCase.EXPECT().Delete("auth_id", input.TodoID)

		rBody, err := json.Marshal(input)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodDelete,
			"/todo",
			bytes.NewBuffer(rBody),
		)
		r = addUser(r)
		r = mux.SetURLVars(r, map[string]string{"todoID": "1"})

		api.Delete(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(rBody)+"\n", w.Body.String())
	})
}

func Test_todoAPI_List(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		api := newTodoAPI(gomock.NewController(t))

		todoUseCase := api.todoUseCase.(*mock_usecase.MockTodoUseCase)
		todoUseCase.EXPECT().List("auth_id", "keyword", "title").
			Return([]*domain.Todo{
				{
					TodoID:  1,
					UserID:  2,
					Title:   "title",
					Content: "content",
					Checked: false,
				},
			}, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			"/todo/list",
			nil,
		)
		r = addUser(r)
		r = mux.SetURLVars(r, map[string]string{"keyword": "keyword", "searchTarget": "title"})

		api.List(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "[{\"TodoID\":1,\"UserID\":2,\"Title\":\"title\",\"Content\":\"content\",\"Checked\":false}]\n", w.Body.String())
	})
}
