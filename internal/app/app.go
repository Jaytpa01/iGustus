package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_igustusDelivery "github.com/Jaytpa01/iGustus/internal/delivery/discord"
	_igustusService "github.com/Jaytpa01/iGustus/internal/service"
	"github.com/Jaytpa01/iGustus/pkg/logger"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func Run() {
	bot, err := initSession(os.Getenv("BOT_TOKEN"))
	if err != nil {
		logger.Log.Error("an error occured create a discord session for BeryBot", zap.Error(err))
		return
	}

	// Cleanly close down the Discord session.
	defer bot.Close()

	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged //discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMessageReactions | discordgo.

	igustusService := _igustusService.NewIgustusService(bot)
	_igustusDelivery.NewDiscordHandler(bot, igustusService)

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		logger.Log.Fatal("error connecting to Discord servers", zap.Error(err))
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	logger.Log.Info("Bot is up and running. Press CTRL-C to shutdown")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func initSession(token string) (*discordgo.Session, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))

	if err != nil {
		return nil, err
	}

	return session, nil
}
