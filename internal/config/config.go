package config

import (
	"fmt"

	"github.com/Jaytpa01/iGustus/internal/entities"
	"github.com/akamensky/argparse"
	"github.com/spf13/viper"
)

type config struct {
	Models map[string]entities.PostRequest
}

var (
	modelsViper *viper.Viper
	Config      config
)

func ReadConfig() error {
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

	for k := range Config.Models {
		Config.Models[k] = configureDefault(Config.Models[k])
	}

	return nil
}

func ConfigureModel(model string, args []string) error {
	if _, ok := Config.Models[model]; !ok {
		return fmt.Errorf("model \"%s\" doesn't exist", model)
	}

	parser := argparse.NewParser("iGustus", "Jay's freakin poggers discordge bot")

	temperature := parser.String("t", "temperature", &argparse.Options{Required: false, Help: "The randomness of the model."})
	tokens := parser.Int("m", "maxtokens", &argparse.Options{Required: false, Help: "Maximum number of tokens the API will return."})

	err := parser.Parse(args)
	if err != nil {
		return fmt.Errorf("error parsing args: %s", err.Error())
	}
	return nil
}

func configureDefault(req entities.PostRequest) entities.PostRequest {
	if req.Temperature == 0 {
		req.Temperature = 0.9
	}

	if req.MaxTokens == 0 {
		req.MaxTokens = 38
	}

	if req.PresencePenalty == 0 {
		req.PresencePenalty = -0.5
	}

	if req.FrequencyPenalty == 0 {
		req.FrequencyPenalty = 2
	}

	return req
}
