package entities

import "time"

type PostRequest struct {
	ChannelID   string
	Args        []string
	OpenAIModel string
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
}
