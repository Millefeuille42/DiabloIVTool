package commands

import (
	"bot/pkg/models"
	"bot/pkg/redisCache"
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
		interactionSendError(s, i, "Error fetching upcoming helltides", 0)
		return
	}

	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltides (invalid location)", 0)
		return
	}

	helltides, err := redisCache.GetUpcomingHelltides()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltides", 0)
		return
	}

	message := ":warning: **This info might might be inaccurate** :warning:\n\n"

	for _, helltide := range helltides.HelltideEvents {
		date := helltide.In(loc).Format(time.RFC850)
		message = fmt.Sprintf("%s%s\n", message, date)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Upcoming helltides:\n```%s```", message),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
