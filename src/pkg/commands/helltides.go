package commands

import (
	"diablo_iv_tool/pkg/fetchers"
	"diablo_iv_tool/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingHelltidesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltides")
		return
	}

	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltides (invalid location)")
		return
	}

	helltides, err := fetchers.GetUpcomingHelltides(loc)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltides")
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Upcoming helltides:\n```%s```", helltides),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
