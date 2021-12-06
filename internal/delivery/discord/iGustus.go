package discord

import (
	"os"
	"strings"

	"github.com/Jaytpa01/iGustus/internal/entities"
	"github.com/bwmarrin/discordgo"
)

func NewDiscordHandler(session *discordgo.Session, service entities.IGustusService) {
	handler := newDiscordDeliveryHandler(service)

	session.AddHandler(handler.CommandsHandler)
}

func (d *discordHandler) CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmdPrefix := os.Getenv("COMMAND_PREFIX")

	if !strings.HasPrefix(m.Content, cmdPrefix) {
		return
	}

	args := strings.Split(m.Content, " ")
	// trim the prefix off of the command. i.e if the prefix was "g!", then "g!pause" will become "pause"
	args[0] = args[0][len(cmdPrefix):]

	// m.ChannelID

	switch args[0] {
	case "test":
		userIDs := []string{}
		for _, user := range m.Mentions {
			userIDs = append(userIDs, user.ID)
		}

		// d.IgustusService.Test(entities.TestRequest{Args: args})

	case "scrape":
		if len(m.Mentions) == 0 {
			s.ChannelMessageSend(m.ChannelID, "Failed to scrape messages. No mentiond users to scrape")
			return
		}

		userIDs := []string{}
		for _, user := range m.Mentions {
			userIDs = append(userIDs, user.ID)
		}

		req := entities.ScrapeRequest{
			ChannelID: m.ChannelID,
			UserIDs:   userIDs,
			Args:      args,
		}

		d.IgustusService.Scrape(req)

	case "post":
		postReq := entities.PostRequest{
			ChannelID: m.ChannelID,
			Args:      args,
		}
		d.IgustusService.Post(postReq)
	}

}
