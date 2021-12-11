package discord

import (
	"os"
	"strings"
	"time"

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
		randReplyReq := entities.RandomReplyRequest{
			ChannelID:     m.ChannelID,
			UserIDToReply: m.Author.ID,
			Timestamp:     time.Now(),
			MsgContent:    m.Content,
		}

		d.IgustusService.RandomlyReply(randReplyReq)
		return
	}

	args := strings.Split(m.Content, " ")
	// trim the prefix off of the command. i.e if the prefix was "g!", then "g!pause" will become "pause"
	args[0] = args[0][len(cmdPrefix):]

	command := strings.ToLower(args[0])

	switch command {
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
			ChannelID:   m.ChannelID,
			Args:        args,
			OpenAIModel: os.Getenv("OPENAI_MODEL_IGUSTUS"),
			Tokens:      38,
		}
		d.IgustusService.Post(postReq)

	case "jiz":
		jizReq := entities.PostRequest{
			ChannelID:   m.ChannelID,
			Args:        args,
			OpenAIModel: os.Getenv("OPENAI_MODEL_JIZ"),
			Tokens:      76,
		}
		d.IgustusService.Post(jizReq)

	case "zep":
		zepReq := entities.PostRequest{
			ChannelID:   m.ChannelID,
			Args:        args,
			OpenAIModel: os.Getenv("OPENAI_MODEL_ZEP"),
			Tokens:      90,
		}
		d.IgustusService.Post(zepReq)

	case "zep2":
		zep2Req := entities.PostRequest{
			ChannelID:   m.ChannelID,
			Args:        args,
			OpenAIModel: os.Getenv("OPENAI_MODEL_ZEP2"),
			APIKey:      os.Getenv("OPENAI_TOKEN_ZEP"),
			Tokens:      40,
		}
		d.IgustusService.Post(zep2Req)

	case "jizus":
		jizusReq := entities.PostRequest{
			ChannelID:   m.ChannelID,
			Args:        args,
			OpenAIModel: os.Getenv("OPENAI_MODEL_JIZUS"),
			Tokens:      38,
		}
		d.IgustusService.Post(jizusReq)

	case "trump":
		trumpReq := entities.PostRequest{
			ChannelID:   m.ChannelID,
			Args:        args,
			OpenAIModel: os.Getenv("OPENAI_MODEL_TRUMP"),
			Tokens:      42,
		}
		d.IgustusService.Post(trumpReq)
	}

}
