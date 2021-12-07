package service

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Jaytpa01/iGustus/internal/entities"
	"github.com/Jaytpa01/iGustus/pkg/emote"
	"github.com/Jaytpa01/iGustus/pkg/logger"
	"github.com/akamensky/argparse"
	"github.com/bwmarrin/discordgo"
	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

type igustusService struct {
	discordSession *discordgo.Session
}

func NewIgustusService(session *discordgo.Session) entities.IGustusService {
	return &igustusService{
		discordSession: session,
	}
}

func (is *igustusService) Test(req entities.TestRequest) {
	// Create new parser object
	parser := argparse.NewParser("iGustus", "Prints provided string to stdout")

	// Create string flag
	s := parser.String("t", "test", &argparse.Options{Required: true, Help: "String to print"})
	// Parse input

	fmt.Println(req.Args)

	err := parser.Parse(req.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Finally print the collected string
	fmt.Println(*s)
}

func (is *igustusService) Post(postReq entities.PostRequest) {
	prompt := ""

	if len(postReq.Args) > 1 {
		prompt = strings.Join(postReq.Args[1:], " ")
	}

	resp, err := createCompletionWithFineTunedModel(prompt, postReq.OpenAIModel)
	if err != nil {
		logger.Log.Error("error posting...", zap.Error(err))
		return
	}

	// if a prompt was provided, make sure the bot actually posts it
	responseText := resp.Choices[0].Text
	if prompt != "" {
		responseText = prompt + responseText
	}

	switch postReq.OpenAIModel {
	case os.Getenv("OPENAI_MODEL_JIZ"):
		responseText += fmt.Sprintf(" - %s", emote.EMOTE_JIZ)

	case os.Getenv("OPENAI_MODEL_IGUSTUS"):
		responseText += fmt.Sprintf(" - %s", emote.EMOTE_FRIGACHAD)

	default:
		responseText += " - wise unknown robot"
	}

	_, err = is.discordSession.ChannelMessageSend(postReq.ChannelID, responseText)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("error posting message to channel with ID: %s", postReq.ChannelID))
		return
	}

}

func createCompletionWithFineTunedModel(prompt, model string) (gogpt.CompletionResponse, error) {
	c := gogpt.NewClient(os.Getenv("OPENAI_TOKEN"))
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Prompt:           prompt,
		Temperature:      0.7,
		Model:            &model,
		MaxTokens:        32,
		PresencePenalty:  -0.5,
		FrequencyPenalty: 0.3,
	}

	return c.CreateCompletionWithFineTunedModel(ctx, req)
}

func (is *igustusService) Scrape(req entities.ScrapeRequest) {
	// Create new parser object
	parser := argparse.NewParser("iGustus", "Prints provided string to stdout")

	// Create string flag
	filename := parser.String("f", "file", &argparse.Options{Required: false, Help: "Filename for csv file of messages"})
	// Parse input
	err := parser.Parse(req.Args)
	if err != nil {
		// is.discordSession.ChannelMessageSend(req.ChannelID, fmt.Sprintf("error parsing arguments:\n %s", parser.Usage(err)))
		// fmt.Println(parser.Usage(err))
	}

	// status message
	msg, _ := is.discordSession.ChannelMessageSend(req.ChannelID, "Scraping messages now...")

	csv := "prompt,completion\n"
	returnedCSV, msgLen, lastMessageID := scrape(is.discordSession, req.ChannelID, "", req.UserIDs)

	if msgLen < 100 {
		logger.Log.Error("msgLen should be 100")
		is.discordSession.ChannelMessageSend(req.ChannelID, "error scraping messages, not enough messages")
		return
	}

	csv += returnedCSV

	for {

		returnedCSV, msgLen, lastMessageID = scrape(is.discordSession, req.ChannelID, lastMessageID, req.UserIDs)
		csv += returnedCSV

		if msgLen < 100 {
			break
		}

	}

	is.discordSession.ChannelMessageEdit(req.ChannelID, msg.ID, "Finsihed scraping, writing to file...")

	time := time.Now().Format("2006-01-02T15-04-05")

	f, err := os.Create(fmt.Sprintf("./junk/%s-%s-%s.csv", *filename, req.ChannelID, time))
	if err != nil {
		logger.Log.Error("error creating output file", zap.Error(err))
		is.discordSession.ChannelMessageEdit(req.ChannelID, msg.ID, fmt.Sprintf("Error creating output file: %s", err.Error()))
		return
	}

	defer f.Close()

	_, err = f.WriteString(csv)
	if err != nil {
		logger.Log.Error("error writing to file", zap.Error(err))
		is.discordSession.ChannelMessageEdit(req.ChannelID, msg.ID, fmt.Sprintf("Error writing messages to file: %s", err.Error()))
		return
	}

	is.discordSession.ChannelMessageEdit(req.ChannelID, msg.ID, "Bing: Finsihed scraping messages")
}

func scrape(s *discordgo.Session, channelID, lastMessageID string, userIdsToScrape []string) (string, int, string) {
	msgs, err := s.ChannelMessages(channelID, 100, lastMessageID, "", "")
	if err != nil {
		logger.Log.Error("error scraping messages", zap.Error(err))
		s.ChannelMessageSend(channelID, fmt.Sprintf("Error scraping messages: %s", err.Error()))
	}

	csv := ""

	for _, msg := range msgs {

		if msg.Content == "" || strings.Contains(msg.Content, "<") || strings.Contains(msg.Content, ">") {
			continue
		}

		for _, id := range userIdsToScrape {
			if msg.Author.ID == id {
				m := strings.ReplaceAll(msg.Content, ",", "")
				m = strings.ReplaceAll(m, "\n", " ")
				m += "\n"
				csv += fmt.Sprintf(",%s", m)
				break
			}
		}

	}
	lastMsgID := ""
	lenMsgs := len(msgs)
	if lenMsgs == 100 {
		lastMsgID = msgs[99].ID
	}

	return csv, len(msgs), lastMsgID
}
