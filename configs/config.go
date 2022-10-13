package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	dbPort     string `mapstructure:"DB_PORT"`
	dbHost     string `mapstructure:"DB_HOST"`
	dbUserName string `mapstructure:"DB_USERNAME"`
	dbPassword string `mapstructure:"DB_PASSWORD"`
	dbName     string `mapstructure: "DB_NAME"`
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
