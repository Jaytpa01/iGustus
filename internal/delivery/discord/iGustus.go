package discord

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/Jaytpa01/iGustus/internal/config"
	"github.com/Jaytpa01/iGustus/internal/entities"
	"github.com/Jaytpa01/iGustus/pkg/util"
	"github.com/bwmarrin/discordgo"
)

func NewDiscordHandler(session *discordgo.Session, service entities.IGustusService, mudService entities.MudService) {
	handler := newDiscordDeliveryHandler(service, mudService)

	session.AddHandler(handler.CommandsHandler)
}

func (d *discordHandler) CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
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

	// if the command matches a model in the models config
	if _, ok := config.Config.Models[command]; ok {
		// if there are arguments, we assume the user wants to configure
		if util.HasArgumentFlag(m.Content) /*|| strings.Contains(m.Content, "-h") || strings.Contains(m.Content, "--help")*/ {
			modelInfo, helpMsg, err := config.ConfigureModel(command, args)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error editing model: %s", err.Error()))
			}

			if helpMsg != "" {
				s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Bot Usage",
							Value: helpMsg,
						},
					},
				})

				return
			}

			if modelInfo != "" {
				s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  fmt.Sprintf("%s config.", command),
							Value: modelInfo,
						},
					},
				})
				return
			}

			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  fmt.Sprintf("%s successfully updated.", command),
						Value: config.Config.Models[command].String(),
					},
				},
			})
			return
		}
		// otherwise, we assume the user just wants to post
		tmp := config.Config.Models[command]
		tmp.ChannelID = m.ChannelID
		tmp.Args = args
		config.Config.Models[command] = tmp

		d.IgustusService.Post(config.Config.Models[command])
		return

	}

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

	case "bots":
		bots := make([]string, 0, len(config.Config.Models))
		for k := range config.Config.Models {
			bots = append(bots, k)
		}

		sort.Strings(bots)

		botList := ""
		for _, bot := range bots {
			botList += bot + "\n"
		}

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Bot List",
					Value: botList,
				},
			},
		})

	case "roll":
		rollReq := entities.RollRequest{
			Content: m.Content,
		}
		rolls, err := d.MudService.Roll(rollReq)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("An error occured trying to roll your dice: %s", err.Error()))
		}

		rollStr := ""
		for _, roll := range rolls {
			rollStr += fmt.Sprintf("Total: %d. Rolls: ", roll.Total)
			for _, r := range roll.Results {
				rollStr += fmt.Sprintf("[%d/%d] ", r, roll.Faces)
			}
			rollStr += "\n"
		}

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Rolls",
					Value: rollStr,
				},
			},
		})

	case "bitch":
		msg := "<@337696913325293578> "
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := r.Intn(180) + 1
		for i := 0; i < n; i++ {
			msg += "bitch "
		}

		s.ChannelMessageSend(m.ChannelID, msg)

	}

}
