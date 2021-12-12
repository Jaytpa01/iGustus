package config

import (
	"github.com/Jaytpa01/iGustus/internal/entities"
	"github.com/spf13/viper"
)

type config struct {
	Models map[string]entities.PostRequest
}

var (
	modelsViper *viper.Viper
	Config      config
)

func readConfig() error {
	modelsViper = viper.New()
	modelsViper.SetConfigName("config")
	modelsViper.SetConfigType("json")
	modelsViper.AddConfigPath("../../")

	err := modelsViper.ReadInConfig()
	if err != nil {
		return err
	}

	err = modelsViper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	return nil
}
