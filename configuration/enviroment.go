package configuration

import (
	"fmt"
	"github.com/spf13/viper"
)

type EnvConfigModel struct {
	DBHost              string `mapstructure:"POSTGRES_HOST"`
	DBUserName          string `mapstructure:"POSTGRES_USER"`
	DBUserPassword      string `mapstructure:"POSTGRES_PASSWORD"`
	DBName              string `mapstructure:"POSTGRES_DB"`
	DBPort              string `mapstructure:"POSTGRES_PORT"`
	SSLMode             string `mapstructure:"SSL_MODE"`
	Port                string `mapstructure:"PORT"`
	STRATZToken         string `mapstructure:"STRATZ_TOKEN"`
	SteamLoginUsernames string `mapstructure:"STEAM_LOGIN_USERNAMES"`
	SteamLoginPasswords string `mapstructure:"STEAM_LOGIN_PASSWORDS"`
}

var EnvConfig EnvConfigModel

func LoadConfig(filePath string) (err error) {
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()

	if viper.ReadInConfig() != nil {
		return
	}

	if err = viper.Unmarshal(&EnvConfig); err != nil {
		return err
	}
	test := EnvConfig
	fmt.Println(test)
	return nil
}
