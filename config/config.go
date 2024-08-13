package config

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	EmailServer string `mapstructure:"EMAIL_SERVER"`
	EmailPort   int    `mapstructure:"EMAIL_PORT"`
	EmailUser   string `mapstructure:"EMAIL_USER"`
	EmailPass   string `mapstructure:"EMAIL_PASS"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}
