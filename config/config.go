package config

import (
	"github.com/spf13/viper"
)

var Config config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("./config/local")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Config = config{
		Profile: viper.GetString("profile"),
		OpenAI: openAI{
			ApiKey:     viper.GetString("openai.api-key"),
			MaxRetries: viper.GetInt("openai.max-retries"),
		},
	}
}

type config struct {
	Profile string
	OpenAI  openAI
}

type openAI struct {
	ApiKey     string
	MaxRetries int
	Models     openAIModels
}

type openAIModels struct {
	Completion []openAIModel
	Embedding  []openAIModel
}

type openAIModel struct{}
