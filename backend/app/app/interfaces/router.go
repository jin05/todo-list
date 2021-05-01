package interfaces

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-list/app/config"
	"todo-list/app/interfaces/api"
	"todo-list/app/interfaces/middleware"
)

func Dispatch(config *config.Config, middlewares middleware.Middlewares, api *api.API) error {
	router := mux.NewRouter().StrictSlash(true)
	for _, mw := range middlewares.List() {
		router.Use(mw)
	}

	router.HandleFunc("/", rootPage).Methods("GET")
	router.HandleFunc("/user", api.UserApi.Signup).Methods("POST")

	return http.ListenAndServe(fmt.Sprintf(":%s", config.Server.ListenPort), router)
}

func rootPage(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "success"); err != nil {
		log.Println(err)
	}
}
