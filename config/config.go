package config

import (
	"github.com/spf13/viper"
)

var Config *config

type config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	DSN string
}

func LoadConfig() *config {
	cfgPath := viper.GetString("config")
	viper.SetConfigName(cfgPath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg *config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	Config = cfg

	return cfg
}
