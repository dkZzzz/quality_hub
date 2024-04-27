package chat

import (
	"context"

	next_api "github.com/dkZzzz/openai_next_api"
	"github.com/dkZzzz/quality_hub/config"
)

const (
	init_question string = `Now I use SonarQube Scan my project, 
	after that, I gain some issue to fix. 
	I want to know how to fix them. 
	Next I will give the code and problem description, 
	and you tell me how to fix it`
)

func Chat(code, message string) (string, error) {
	client := next_api.NewClient(config.Cfg.OpenaiSK)
	req := next_api.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []next_api.ChatCompletionMessage{
			{
				Role:    "user",
				Content: init_question + "\n" + "my code is: " + code + "\n" + "the problem is: " + message + "\n" + "how to fix it?",
			},
		},
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
