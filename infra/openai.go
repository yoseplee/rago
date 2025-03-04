package infra

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/yoseplee/rago/config"
)

var OpenAIClient *openai.Client

var LinecorpOpenAIClient *openai.Client

func init() {
	OpenAIClient = openai.NewClient(
		option.WithBaseURL(config.Config.OpenAI["public"].BaseUrl),
		option.WithAPIKey(config.Config.OpenAI["public"].ApiKey),
		option.WithMaxRetries(config.Config.OpenAI["public"].MaxRetries),
	)

	LinecorpOpenAIClient = openai.NewClient(
		option.WithBaseURL(config.Config.OpenAI["linecorp"].BaseUrl),
		option.WithAPIKey(config.Config.OpenAI["linecorp"].ApiKey),
		option.WithMaxRetries(config.Config.OpenAI["linecorp"].MaxRetries),
	)
}
