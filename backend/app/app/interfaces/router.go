package interfaces

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-list/app/config"
	"todo-list/app/interfaces/api"
	"todo-list/app/interfaces/middleware"
)

var (
	muxAdapter *gorillamux.GorillaMuxAdapter
)

func Dispatch(config *config.Config, middlewares middleware.Middlewares, api *api.API) error {

	if config.Server.IsProduction {
		router := mux.NewRouter()
		for _, mw := range middlewares.List() {
			router.Use(mw)
		}
		setRouter(router, api)
		muxAdapter = gorillamux.New(router)
		lambda.Start(lambdaHandler)
		return nil
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, mw := range middlewares.List() {
		router.Use(mw)
	}

	setRouter(router, api)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Server.ListenPort), router)
}

func setRouter(router *mux.Router, api *api.API) {
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/user", api.UserApi.Signup).Methods("POST")
}

func rootPage(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "success"); err != nil {
		log.Println(err)
	}
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res, err := muxAdapter.Proxy(req)
	if err != nil {
		log.Println(err)
	}
	return res, err
}
