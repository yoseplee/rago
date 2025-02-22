package infra

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/yoseplee/rago/config"
)

var OpenAIClient *openai.Client

func init() {
	OpenAIClient = openai.NewClient(
		option.WithAPIKey(config.Config.ApiKey),
		option.WithMaxRetries(config.Config.MaxRetries),
	)
}
