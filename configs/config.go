package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUser 		string `mapstructure:"DB_USER"`
	DBPass 		string `mapstructure:"DB_PASSWORD"`
	DBName 		string `mapstructure:"DB_NAME"`
	DBHost 		string `mapstructure:"DB_HOST"`
	DBPort		string `mapstructure:"DB_PORT"`
	HttpPort        string `mapstructure:"HTTP_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}





