package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"todo-list/app/interfaces/middleware"
	"todo-list/app/usecase"
)

type createInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type updateInput struct {
	TodoID  int64  `json:"todoID"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Checked bool   `json:"checked"`
}

type deleteInput struct {
	TodoID int64 `json:"todoID"`
}

type TodoAPI interface {
	Handler(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type todoAPI struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoAPI(todoUseCase usecase.TodoUseCase) TodoAPI {
	return &todoAPI{todoUseCase: todoUseCase}
}

func (a *todoAPI) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return
	case http.MethodGet:
		a.Get(w, r)
	case http.MethodPost:
		a.Create(w, r)
	case http.MethodPut:
		a.Update(w, r)
	case http.MethodDelete:
		a.Delete(w, r)
	}
}

func (a *todoAPI) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザ情報を取得
	cUser := middleware.UserForContext(ctx)

	vars := mux.Vars(r)
	todoID, err := strconv.ParseInt(vars["todoID"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := a.todoUseCase.Get(cUser.AuthID, todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *todoAPI) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザ情報を取得
	cUser := middleware.UserForContext(ctx)

	input := createInput{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := a.todoUseCase.Create(cUser.AuthID, input.Title, input.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *todoAPI) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザ情報を取得
	cUser := middleware.UserForContext(ctx)

	input := updateInput{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := a.todoUseCase.Update(cUser.AuthID, input.TodoID, input.Title, input.Content, input.Checked)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *todoAPI) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザ情報を取得
	cUser := middleware.UserForContext(ctx)

	input := deleteInput{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.todoUseCase.Delete(cUser.AuthID, input.TodoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *todoAPI) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザ情報を取得
	cUser := middleware.UserForContext(ctx)

	vars := mux.Vars(r)
	keyWard := vars["keyWard"]

	todoList, err := a.todoUseCase.List(cUser.AuthID, keyWard)
	if err != nil {
		return
	}

	if err = json.NewEncoder(w).Encode(todoList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
