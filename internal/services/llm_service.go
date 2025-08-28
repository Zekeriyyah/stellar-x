package services

import (
	"context"
	"fmt"
	"os"

	"github.com/Zekeriyyah/stellar-x/pkg"
	"github.com/sashabaranov/go-openai"
)

type AIService struct {
	client      *openai.Client
	fxService   *FXService
}

func NewAIService(fxService *FXService) *AIService {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil
	}

	client := openai.NewClient(apiKey)
	return &AIService{
		client:    client,
		fxService: fxService,
	}
}

// Ask: queries the LLM with FX context
func (s *AIService) Ask(query string) (string, error) {

	// Scrape query for currencies if present
	currencies := pkg.ScrapeQuery(query)

	// Get real FX data to ground the response
	var rate float64
	if len(currencies) == 2 {
		rate, _ = s.fxService.GetRate(currencies[0], currencies[1])

	} 

	systemPrompt := `
You are FX AI Assistant, a financial expert for the StellarX payment platform.
Answer concisely and accurately.
Use this real FX rate: 1 cNGN = ` + fmt.Sprintf("%.6f", rate) + ` USDx
If asked about trends, say "Based on recent data" and keep it factual.
Never make up rates.`
	
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	// ✅ Safe: Check if choices exist
	if len(resp.Choices) == 0 {
		return "No response from AI", nil
	}

	return resp.Choices[0].Message.Content, nil
}