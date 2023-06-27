package commands

import (
	"bot/pkg/models"
	"bot/pkg/redisCache"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingHelltideCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helltide, err := redisCache.GetHelltide()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltide", 0)
		return
	}

	guild := models.GuildModel{}
	_, err = guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss", 0)
		return
	}

	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss (invalid location)", 0)
		return
	}

	untilHelltide := time.Until(helltide.StartTime)
	date := helltide.StartTime.In(loc).Format(time.TimeOnly)

	minutes := int(untilHelltide.Minutes())

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Next helltide in %02dm (%s)", minutes, date),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
