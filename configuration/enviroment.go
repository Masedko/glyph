package configuration

import (
	"github.com/spf13/viper"
	"os"
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
	EnvConfig.DBHost = os.Getenv("POSTGRES_HOST")
	EnvConfig.DBUserName = os.Getenv("POSTGRES_USER")
	EnvConfig.DBUserPassword = os.Getenv("POSTGRES_PASSWORD")
	EnvConfig.DBName = os.Getenv("POSTGRES_DB")
	EnvConfig.DBPort = os.Getenv("POSTGRES_PORT")
	EnvConfig.SSLMode = os.Getenv("SSL_MODE")
	EnvConfig.Port = os.Getenv("PORT")
	EnvConfig.STRATZToken = os.Getenv("STRATZ_TOKEN")
	EnvConfig.SteamLoginUsernames = os.Getenv("STEAM_LOGIN_USERNAMES")
	EnvConfig.SteamLoginPasswords = os.Getenv("STEAM_LOGIN_PASSWORDS")

	if viper.ReadInConfig() != nil {
		return
	}

	if err = viper.Unmarshal(&EnvConfig); err != nil {
		return err
	}
	return nil
}
