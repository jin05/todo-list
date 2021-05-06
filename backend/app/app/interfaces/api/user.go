package api

import (
	"encoding/json"
	"net/http"
	"todo-list/app/interfaces/middleware"
	"todo-list/app/usecase"
)

type UserAPI interface {
	Signup(w http.ResponseWriter, r *http.Request)
}

type userAPI struct {
	userUseCase usecase.UserUseCase
}

func NewUserAPI(userUseCase usecase.UserUseCase) UserAPI {
	return &userAPI{userUseCase: userUseCase}
}

func (a *userAPI) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザ情報を取得
	cUser := middleware.UserForContext(ctx)

	user, err := a.userUseCase.Save(cUser.AuthID, cUser.Name, cUser.Email)
	if err != nil {
		middleware.SetError(ctx, err)
		return
	}

	if err = json.NewEncoder(w).Encode(user); err != nil {
		middleware.SetError(ctx, err)
	}
}
