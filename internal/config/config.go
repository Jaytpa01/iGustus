package config

import (
	"fmt"
	"strings"

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
	modelsViper.AddConfigPath(".")
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

func ConfigureModel(model string, args []string) (string, string, error) {
	if _, ok := Config.Models[model]; !ok {
		return "", "", fmt.Errorf("model \"%s\" doesn't exist", model)
	}

	parser := argparse.NewParser("iGustus", "Jay's freakin poggers discordge bot")
	parser.DisableHelp()

	temperature := parser.Float("t", "temperature", &argparse.Options{Required: false, Help: "The randomness of the model. Between 0 and 1.0", Default: -3.0})
	tokens := parser.Int("m", "maxtokens", &argparse.Options{Required: false, Help: "Maximum number of tokens the API will return."})
	freqPenalty := parser.Float("f", "frequency", &argparse.Options{Required: false, Help: "Set the frequency penalty of the model. Accepts values between -2.0 and 2.0", Default: -3.0})
	presPenalty := parser.Float("p", "presence", &argparse.Options{Required: false, Help: "Set the presence penalty of the model. Between -2.0 and 2.0.", Default: -3.0})
	signature := parser.String("s", "signature", &argparse.Options{Required: false, Help: "Set the model's signature."})
	help := parser.Flag("h", "help", &argparse.Options{Required: false})
	info := parser.Flag("i", "info", &argparse.Options{Required: false, Help: "Returns current configuration of the model."})

	err := parser.Parse(args)
	if err != nil {
		// only return error if its not an unknown argument error
		if !strings.Contains(err.Error(), "unknown arguments") {
			return "", "", fmt.Errorf("error parsing args: %s", err.Error())
		}
	}

	if *help {
		return "", parser.Usage(""), nil
	}

	if *info {
		return Config.Models[model].String(), "", nil
	}

	if *signature != "" {
		modelsViper.Set(fmt.Sprintf("models.%s.signature", model), *signature)
	}

	// if temp value is given
	if *temperature != -3.0 {
		// clamp values
		if *temperature < 0 {
			*temperature = 0
		} else if *temperature > 1 {
			*temperature = 1
		}
		modelsViper.Set(fmt.Sprintf("models.%s.temperature", model), *temperature)
	}

	if *freqPenalty != -3.0 {
		// clamp
		if *freqPenalty < -2.0 {
			*freqPenalty = -2
		} else if *freqPenalty > 2 {
			*freqPenalty = 2.0
		}
		modelsViper.Set(fmt.Sprintf("models.%s.frequencypenalty", model), *freqPenalty)
	}

	if *presPenalty != -3.0 {
		if *presPenalty < -2 {
			*presPenalty = -2
		} else if *presPenalty > 2 {
			*presPenalty = 2
		}
		modelsViper.Set(fmt.Sprintf("models.%s.presencepenalty", model), *presPenalty)
	}

	if *tokens != 0 {
		if *tokens < 5 {
			*tokens = 5
		}
		modelsViper.Set(fmt.Sprintf("models.%s.maxtokens", model), *tokens)
	}

	err = modelsViper.WriteConfig()
	if err != nil {
		return "", "", err
	}

	err = modelsViper.Unmarshal(&Config)
	if err != nil {
		return "", "", err
	}

	return "", "", nil
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

	if req.Signature == "" {
		req.Signature = "wise unknown robot"
	}

	return req
}
