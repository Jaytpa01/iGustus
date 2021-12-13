package entities

import (
	"fmt"
	"time"
)

type PostRequest struct {
	ChannelID   string
	Args        []string
	OpenAIModel string
	APIKey      string
	Signature   string

	MaxTokens        int
	Temperature      float32
	PresencePenalty  float32
	FrequencyPenalty float32
}

func (p PostRequest) String() string {
	return fmt.Sprintf("Max Token: %d\nTemperature: %.2f\nPresence Penalty: %.2f\nFrequency Penalty: %.2f\nSignature: %s\n", p.MaxTokens, p.Temperature, p.PresencePenalty, p.FrequencyPenalty, p.Signature)
}

type ScrapeRequest struct {
	ChannelID string
	UserIDs   []string
	Args      []string
}

type TestRequest struct {
	ChannelID string
	Args      []string
}

type RandomReplyRequest struct {
	ChannelID     string
	UserIDToReply string
	Timestamp     time.Time
	MsgContent    string
}
