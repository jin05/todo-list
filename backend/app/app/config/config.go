package config

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
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
	Port   int
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

	DBUser       string `envconfig:"DB_USER"`
	DBHost       string `envconfig:"DB_HOST"`
	DBPort       int    `envconfig:"DB_PORT"`
	DBPassword   string `envconfig:"DB_PASSWORD"`
	DBName       string `envconfig:"DB_NAME"`
	DBSecretName string `envconfig:"DB_SECRET_NAME"`

	AwsUserPoolId       string `envconfig:"AWS_USER_POOL_ID" required:"true"`
	AwsUserPoolClientId string `envconfig:"AWS_USER_POOL_CLIENT_ID" required:"true"`
}

type DatabaseInfo struct {
	UserName             string `json:"username"`
	Password             string `json:"password"`
	Engine               string `json:"engine"`
	Host                 string `json:"host"`
	Port                 int    `json:"port"`
	DBName               string `json:"dbname"`
	DBInstanceIdentifier string `json:"dbInstanceIdentifier"`
	Proxy                string `json:"proxy"`
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

	awsConfig := &aws.Config{
		Region: aws.String("ap-northeast-1"),
	}

	conf.AWS = AWSConfig{
		Config:           awsConfig,
		UserPoolID:       env.AwsUserPoolId,
		UserPoolClientID: env.AwsUserPoolClientId,
	}

	if conf.Server.IsProduction {
		dbInfo, err := getDatabaseInfo(&conf, env.DBSecretName)
		if err != nil {
			return nil, err
		}
		conf.DB = DBConfig{
			User:   dbInfo.UserName,
			Pass:   dbInfo.Password,
			Host:   dbInfo.Proxy,
			Port:   dbInfo.Port,
			DBName: dbInfo.DBName,
		}
	} else {
		conf.DB = DBConfig{
			User:   env.DBUser,
			Pass:   env.DBPassword,
			Host:   env.DBHost,
			Port:   env.DBPort,
			DBName: env.DBName,
		}
	}

	return &conf, nil
}

func getDatabaseInfo(config *Config, secretName string) (*DatabaseInfo, error) {
	svc := secretsmanager.New(session.New(), config.AWS.Config)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}
	if result.SecretString == nil {
		return nil, errors.New("secret string is nil.")
	}
	info := &DatabaseInfo{}
	err = json.Unmarshal([]byte(*result.SecretString), info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
