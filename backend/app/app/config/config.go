package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	ListenPort   string
	Environment  string
	IsProduction bool
	AllowOrigin  string
}

type DBConfig struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DBName string
}

type AWSConfig struct {
	Config           *aws.Config
	UserPoolID       string
	UserPoolClientID string
}

type Config struct {
	Server ServerConfig
	DB     DBConfig
	AWS    AWSConfig
}

type envConfig struct {
	ServerListenPort  string `envconfig:"LISTEN_PORT" required:"true"`
	ServerEnvironment string `envconfig:"ENV_NAME" required:"true"`
	ServerAllowOrigin string `envconfig:"ALLOW_ORIGIN" required:"true"`

	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     string `envconfig:"DB_PORT" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`

	AWSRegion           string `envconfig:"AWS_REGION" required:"true"`
	AwsUserPoolId       string `envconfig:"AWS_USER_POOL_ID" required:"true"`
	AwsUserPoolClientId string `envconfig:"AWS_USER_POOL_CLIENT_ID" required:"true"`
}

func NewConfig() (*Config, error) {
	conf := Config{}

	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		return nil, err
	}

	conf.Server = ServerConfig{
		ListenPort:   env.ServerListenPort,
		Environment:  env.ServerEnvironment,
		IsProduction: env.ServerEnvironment == "prod",
		AllowOrigin:  env.ServerAllowOrigin,
	}

	conf.DB = DBConfig{
		User:   env.DBUser,
		Pass:   env.DBPassword,
		Host:   env.DBHost,
		Port:   env.DBPort,
		DBName: env.DBName,
	}

	awsConfig := &aws.Config{
		Region: aws.String(env.AWSRegion),
	}

	conf.AWS = AWSConfig{
		Config:           awsConfig,
		UserPoolID:       env.AwsUserPoolId,
		UserPoolClientID: env.AwsUserPoolClientId,
	}

	return &conf, nil
}
