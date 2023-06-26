package configuration

import (
	"github.com/spf13/viper"
)

type EnvConfigModel struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	SSLMode        string `mapstructure:"SSL_MODE"`
	STRATZToken    string `mapstructure:"STRATZ_TOKEN"`
}

var EnvConfig EnvConfigModel

func LoadConfig(filePath string) (err error) {
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)

	viper.AutomaticEnv()

	if viper.ReadInConfig() != nil {
		return
	}

	return viper.Unmarshal(&EnvConfig)
}
