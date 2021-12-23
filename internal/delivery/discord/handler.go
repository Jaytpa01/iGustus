package discord

import (
	"github.com/Jaytpa01/iGustus/internal/entities"
	"github.com/Jaytpa01/iGustus/internal/entities/mud"
)

type discordHandler struct {
	IgustusService entities.IGustusService
	MudService     mud.MudService
}

func newDiscordDeliveryHandler(service entities.IGustusService, mudService mud.MudService) entities.DiscordHandler {
	return &discordHandler{
		IgustusService: service,
		MudService:     mudService,
	}
}
