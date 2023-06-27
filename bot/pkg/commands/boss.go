package commands

import (
	"bot/pkg/models"
	"bot/pkg/redisCache"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingBossCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	worldBoss, err := redisCache.GetWorldBoss()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss", 0)
		return
	}

	zone, err := redisCache.GetWorldBossZone()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss", 0)
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

	untilBoss := time.Until(worldBoss.StartTime)
	date := worldBoss.StartTime.In(loc).Format(time.TimeOnly)

	hours := int(untilBoss.Hours())
	minutes := int(untilBoss.Minutes()) - (hours * 60)

	if zone == "no-boss" || zone == "" {
		zone = ""
	} else {
		zone = fmt.Sprintf("at %s ", zone)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Next world boss is ***%s*** in %dh%02dm %s(%s)",
				worldBoss.Boss,
				hours,
				minutes,
				zone,
				date),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
