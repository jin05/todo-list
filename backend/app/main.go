package main

import (
	"go.uber.org/dig"
	"log"
	"os"
	"todo-list/app/config"
	"todo-list/app/domain/repository"
	"todo-list/app/infrastructure/database"
	"todo-list/app/interfaces"
	"todo-list/app/interfaces/api"
	"todo-list/app/interfaces/middleware"
	"todo-list/app/usecase"
)

func createDIContainer() *dig.Container {
	container := dig.New()
	container.Provide(middleware.NewMiddlewares)
	container.Provide(middleware.NewAuthMiddleware)
	container.Provide(api.NewAPI)
	container.Provide(api.NewUserAPI)
	container.Provide(usecase.NewUserUseCase)
	container.Provide(repository.NewUserRepository)

	if os.Getenv("ENV_NAME") != "test" {
		container.Provide(config.NewConfig)
		container.Provide(database.NewDB)
	}

	return container
}

func main() {
	if err := createDIContainer().Invoke(func(config *config.Config, middlewares middleware.Middlewares, api *api.API) error {
		return interfaces.Dispatch(config, middlewares, api)
	}); err != nil {
		log.Println(err)
	}
}
