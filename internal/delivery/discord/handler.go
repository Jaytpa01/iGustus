package discord

import "github.com/Jaytpa01/iGustus/internal/entities"

type discordHandler struct {
	IgustusService entities.IGustusService
}

func newDiscordDeliveryHandler(service entities.IGustusService) entities.DiscordHandler {
	return &discordHandler{
		IgustusService: service,
	}
}
