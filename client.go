package main

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"rago/config"
)

type Client struct {
	OpenAIClient *openai.Client
}

func NewClient() Client {
	return Client{
		OpenAIClient: openai.NewClient(
			option.WithAPIKey(config.Config.ApiKey),
			option.WithMaxRetries(config.Config.MaxRetries),
		),
	}
}
