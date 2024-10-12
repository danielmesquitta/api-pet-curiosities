package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"

	"github.com/danielmesquitta/api-pet-curiosities/internal/config"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/gpt"
)

type Client struct {
	openAIClient *openai.Client
}

func NewClient(env *config.Env) *Client {
	client := openai.NewClient(env.OpenAIToken)

	return &Client{
		openAIClient: client,
	}
}

func (o *Client) CreateChatCompletion(message string) (string, error) {
	resp, err := o.openAIClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return "", errs.New(err)
	}

	return resp.Choices[0].Message.Content, nil
}

var _ gpt.Provider = (*Client)(nil)
