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
	// Check if the file exists
	if _, err = os.Stat(filePath); err == nil {
		viper.SetConfigFile(filePath)

		// Attempt to read the configuration file
		if err = viper.ReadInConfig(); err != nil {
			return err // File exists but could not be read
		}
	} else {
		envs := []string{"POSTGRES_HOST", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB", "POSTGRES_PORT",
			"SSL_MODE", "PORT", "STRATZ_TOKEN", "STEAM_LOGIN_USERNAMES", "STEAM_LOGIN_PASSWORDS"}
		for _, env := range envs {
			if err = viper.BindEnv(env); err != nil {
				return err
			}
		}
		viper.AutomaticEnv()
	}

	if err = viper.Unmarshal(&EnvConfig); err != nil {
		return err
	}
	return nil
}
