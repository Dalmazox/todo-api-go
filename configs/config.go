package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

const (
	PORT              = "PORT"
	USE_SWAGGER       = "USE_SWAGGER"
	CONNECTION_STRING = "CONNECTION_STRING"
)

type config struct {
	Server           Server
	ConnectionString string
}

type Server struct {
	Port       string
	UseSwagger bool
}

func init() {
	viper := viper.New()

	viper.SetDefault(PORT, 8080)
	viper.SetDefault(USE_SWAGGER, false)

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	viper.ReadInConfig()

	cfg = &config{
		ConnectionString: viper.GetString(CONNECTION_STRING),
		Server: Server{
			Port:       viper.GetString(PORT),
			UseSwagger: viper.GetBool(USE_SWAGGER),
		},
	}
}

func GetConfig() config {
	return *cfg
}
