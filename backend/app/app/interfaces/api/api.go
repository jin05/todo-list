package api

type API struct {
	UserApi UserAPI
}

func NewAPI(userApi UserAPI) *API {
	return &API{UserApi: userApi}
}
