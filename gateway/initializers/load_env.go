package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT                  string `mapstructure:"PORT"`
	AUTH_SERVICE_HOST     string `mapstructure:"AUTH_SERVICE_HOST"`
	MERCHANT_SERVICE_HOST string `mapstructure:"MERCHANT_SERVICE_HOST"`
	ORDER_SERVICE_HOST    string `mapstructure:"ORDER_SERVICE_HOST"`
	PAYMENT_SERVICE_HOST  string `mapstructure:"PAYMENT_SERVICE_HOST"`
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
