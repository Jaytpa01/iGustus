package entities

import "time"

type PostRequest struct {
	ChannelID   string
	Args        []string
	OpenAIModel string
	APIKey      string
	Tokens      int
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
