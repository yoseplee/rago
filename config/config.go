package config

import (
	"encoding/json"
	"fmt"

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

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	prettyConfig, _ := json.MarshalIndent(Config, "", "  ")
	fmt.Println(string(prettyConfig))

}

type config struct {
	Profile        string                  `yaml:"profile"`
	Ingesters      map[string]ingester     `yaml:"ingesters"`
	Retrievers     map[string]retriever    `yaml:"retrievers"`
	KnowledgeBases knowledgeBasesConfig    `yaml:"knowledgeBases"`
	OpenAI         map[string]openAIConfig `yaml:"openai"`
}

type ingester struct {
	DocumentLoader   documentLoaderConfig   `yaml:"document_loader"`
	Embedding        embeddingConfig        `yaml:"embedding"`
	KnowledgeBaseAdd knowledgeBaseAddConfig `yaml:"knowledgeBaseAdd"`
}

type retriever struct {
	Embedding           embeddingConfig           `yaml:"embedding"`
	KnowledgeBaseSearch knowledgeBaseSearchConfig `yaml:"knowledgeBaseSearch"`
}

type documentLoaderConfig struct {
	Kind     string `yaml:"kind"`
	FilePath string `yaml:"filePath"`
}

type knowledgeBaseAddConfig struct {
	Kind       string `yaml:"kind"`
	Collection string `yaml:"collection"`
}

type knowledgeBaseSearchConfig struct {
	Kind       string `yaml:"kind"`
	Collection string `yaml:"collection"`
	TopK       int    `yaml:"topK"`
}

type knowledgeBasesConfig struct {
	Opensearch map[string]opensearchConfig `yaml:"opensearch"`
}

type opensearchConfig struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type embeddingConfig struct {
	Source    string `yaml:"source"`
	Model     string `yaml:"model"`
	Dimension int    `yaml:"dimension"`
}

type openAIConfig struct {
	BaseUrl    string `yaml:"baseUrl,omitempty"`
	ApiKey     string `yaml:"apiKey"`
	MaxRetries int    `yaml:"maxRetries"`
}
