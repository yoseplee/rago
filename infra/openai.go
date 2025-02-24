package infra

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/yoseplee/rago/config"
)

var OpenAIClient *openai.Client

func init() {
	OpenAIClient = openai.NewClient(
		option.WithBaseURL(config.Config.OpenAI.BaseUrl),
		option.WithAPIKey(config.Config.OpenAI.ApiKey),
		option.WithMaxRetries(config.Config.OpenAI.MaxRetries),
	)
}
