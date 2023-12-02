package configuration

import (
	"github.com/spf13/viper"
	"os"
)

type EnvConfigModel struct {
	DBHost             string `mapstructure:"POSTGRES_HOST"`
	DBUserName         string `mapstructure:"POSTGRES_USER"`
	DBUserPassword     string `mapstructure:"POSTGRES_PASSWORD"`
	DBName             string `mapstructure:"POSTGRES_DB"`
	DBPort             string `mapstructure:"POSTGRES_PORT"`
	SSLMode            string `mapstructure:"SSL_MODE"`
	Port               string `mapstructure:"PORT"`
	STRATZToken        string `mapstructure:"STRATZ_TOKEN"`
	SteamLoginUsername string `mapstructure:"STEAM_LOGIN_USERNAME"`
	SteamLoginPassword string `mapstructure:"STEAM_LOGIN_PASSWORD"`
	SteamTwoFactorCode string `mapstructure:"STEAM_TWO_FACTOR_CODE"`
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
	EnvConfig.SteamLoginUsername = os.Getenv("STEAM_LOGIN_USERNAME")
	EnvConfig.SteamLoginPassword = os.Getenv("STEAM_LOGIN_PASSWORD")
	EnvConfig.SteamTwoFactorCode = os.Getenv("STEAM_TWO_FACTOR_CODE")

	if viper.ReadInConfig() != nil {
		return
	}

	return viper.Unmarshal(&EnvConfig)
}
