package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("./config/local")
	viper.SetConfigType("yaml")
	viper.SetDefault("openai.base-url", "https://api.openai.com/v1/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Config = config{
		Profile: viper.GetString("profile"),
		OpenAI: openAI{
			BaseUrl:    viper.GetString("openai.base-url"),
			ApiKey:     viper.GetString("openai.api-key"),
			MaxRetries: viper.GetInt("openai.max-retries"),
		},
	}
	fmt.Printf("%+v\n", Config)
}

type config struct {
	Profile string
	OpenAI  openAI
}

type openAI struct {
	BaseUrl    string
	ApiKey     string
	MaxRetries int
	Models     openAIModels
}

type openAIModels struct {
	Completion []openAIModel
	Embedding  []openAIModel
}

type openAIModel struct{}
