package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	EnableAutoMigrate string `mapstructure:"ENABLE_AUTO_MIGRATE"`
	DBHost            string `mapstructure:"POSTGRES_HOST"`
	DBUserName        string `mapstructure:"POSTGRES_USER"`
	DBUserPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	DBName            string `mapstructure:"POSTGRES_DB"`
	DBPort            string `mapstructure:"POSTGRES_PORT"`
	ServerPort        string `mapstructure:"PORT"`
	DB_TIMEZONE       string `mapstructure:"DB_TIMEZONE"`
	SSLMode           string `mapstructure:"SSL_MODE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenPublicKey string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`

	RedisHost    string `mapstructure:"REDIS_HOST"`
	ReisPassword string `mapstructure:"REDIS_PASSWORD"`

	AMQP_SERVER_PORT     string `mapstructure:"AMQP_SERVER_PORT"`
	AMQP_SERVER_HOST     string `mapstructure:"AMQP_SERVER_HOST"`
	AMQP_SERVER_USER     string `mapstructure:"AMQP_SERVER_USER"`
	AMQP_SERVER_PASSWORD string `mapstructure:"AMQP_SERVER_PASSWORD"`

	USER_GRPC_SERVER_PORT string `mapstructure:"USER_GRPC_SERVER_PORT"`
	USER_GRPC_SERVER_HOST string `mapstructure:"USER_GRPC_SERVER_HOST"`

	MERCHANT_GRPC_SERVER_PORT string `mapstructure:"MERCHANT_GRPC_SERVER_PORT"`
	MERCHANT_GRPC_SERVER_HOST string `mapstructure:"MERCHANT_GRPC_SERVER_HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
