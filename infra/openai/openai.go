package openai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/yoseplee/rago/config"
)

var OpenAIClient *DefaultClient

var LinecorpOpenAIClient *DefaultClient

type EmbeddingGeneratable interface {
	Embedding(modelName string, documents []string) (*EmbeddingResponse, error)
}

type EmbeddingResponse struct {
	*openai.CreateEmbeddingResponse
}

type Chattable interface {
	Chat(modelName string, prompts []string) (*ChatResponse, error)
}

type Client interface {
	EmbeddingGeneratable
	Chattable
}

func init() {
	OpenAIClient = &DefaultClient{
		Client: openai.NewClient(
			option.WithBaseURL(config.Config.OpenAI["public"].BaseUrl),
			option.WithAPIKey(config.Config.OpenAI["public"].ApiKey),
			option.WithMaxRetries(config.Config.OpenAI["public"].MaxRetries),
		),
	}

	LinecorpOpenAIClient = &DefaultClient{
		Client: openai.NewClient(
			option.WithBaseURL(config.Config.OpenAI["linecorp"].BaseUrl),
			option.WithAPIKey(config.Config.OpenAI["linecorp"].ApiKey),
			option.WithMaxRetries(config.Config.OpenAI["linecorp"].MaxRetries),
		),
	}
}

type DefaultClient struct {
	Client *openai.Client
}

func (d DefaultClient) Embedding(modelName string, documents []string) (*EmbeddingResponse, error) {
	res, err := d.Client.Embeddings.New(
		context.TODO(),
		openai.EmbeddingNewParams{
			Input: openai.F[openai.EmbeddingNewParamsInputUnion](openai.EmbeddingNewParamsInputArrayOfStrings(documents)),
			Model: openai.F(modelName),
		},
	)
	if err != nil {
		return nil, err
	}

	return &EmbeddingResponse{res}, nil
}

type ChatResponse struct {
	*openai.ChatCompletion
}

func (d DefaultClient) Chat(modelName string, prompts []string) (*ChatResponse, error) {
	// TODO: improve prompts - add assist message.

	var messages []openai.ChatCompletionMessageParamUnion
	for _, prompt := range prompts {
		messages = append(messages, openai.UserMessage(prompt))
	}

	chatCompletion, err := d.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F(messages),
		Model:    openai.F(modelName),
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(chatCompletion.Choices[0].Message.Content)
	return &ChatResponse{chatCompletion}, nil
}
