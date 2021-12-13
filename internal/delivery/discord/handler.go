package discord

import "github.com/Jaytpa01/iGustus/internal/entities"

type discordHandler struct {
	IgustusService entities.IGustusService
	MudService     entities.MudService
}

func newDiscordDeliveryHandler(service entities.IGustusService, mudService entities.MudService) entities.DiscordHandler {
	return &discordHandler{
		IgustusService: service,
		MudService:     mudService,
	}
}
