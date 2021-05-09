package test

import (
	"os"
)

var (
	OriginEnvs = map[string]string{
		"ENV_NAME":     "test",
		"LISTEN_PORT":  "1234",
		"ALLOW_ORIGIN": "http://localhost:3000",

		"DB_USER":        "root",
		"DB_HOST":        "127.0.0.1",
		"DB_PORT":        "3306",
		"DB_PASSWORD":    "mysql",
		"DB_NAME":        "todo",
		"DB_SECRET_NAME": "local",

		"AWS_USER_POOL_ID":        "aws_user_pool_id",
		"AWS_USER_POOL_CLIENT_ID": "aws_user_pool_client_id",
	}
)

func MakeTestEnvs() map[string]string {
	envs := map[string]string{}
	for key, value := range OriginEnvs {
		envs[key] = value
	}

	return envs
}

func SetupEnvs(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}

func ClearEnvs(envs map[string]string) {
	for key, _ := range envs {
		os.Unsetenv(key)
	}
}
