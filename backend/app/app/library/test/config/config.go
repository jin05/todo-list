package test_config

import (
	"todo-list/app/config"
	"todo-list/app/library/test"
)

func NewTestConfig() *config.Config {
	test.SetupEnvs(test.OriginEnvs)
	defer test.ClearEnvs(test.OriginEnvs)

	conf, _ := config.NewConfig()
	return conf
}
