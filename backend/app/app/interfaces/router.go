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
	router := mux.NewRouter().StrictSlash(true)
	for _, mw := range middlewares.List() {
		router.Use(mw)
	}

	router.HandleFunc("/", rootPage).Methods("GET")
	router.HandleFunc("/user", api.UserApi.Signup).Methods("POST")

	if config.Server.IsProduction {
		muxAdapter = gorillamux.New(router)
		lambda.Start(lambdaHandler)
		return nil
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", config.Server.ListenPort), router)
}

func rootPage(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "success"); err != nil {
		log.Println(err)
	}
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req.HTTPMethod = "POST"
	res, err := muxAdapter.Proxy(req)
	if err != nil {
		log.Println(err)
	}
	return res, err
}
