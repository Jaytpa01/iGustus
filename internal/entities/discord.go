package entities

import "github.com/bwmarrin/discordgo"

type DiscordHandler interface {
	CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate)
}

type IGustusService interface {
	Test(TestRequest)
	Scrape(ScrapeRequest)
	Post(PostRequest)
}
