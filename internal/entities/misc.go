package entities

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
