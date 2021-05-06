package api

type API struct {
	UserApi UserAPI
	TodoAPI TodoAPI
}

func NewAPI(userApi UserAPI, todoAPI TodoAPI) *API {
	return &API{UserApi: userApi, TodoAPI: todoAPI}
}
