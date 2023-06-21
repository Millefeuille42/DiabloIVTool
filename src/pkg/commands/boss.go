package commands

import (
	"diablo_iv_tool/pkg/fetchers"
	"diablo_iv_tool/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingBossCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	upcomingBoss, err := fetchers.GetUpcomingBossData()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss")
		return
	}

	guild := models.GuildModel{}
	_, err = guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss")
		return
	}

	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss (invalid location)")
		return
	}

	name, hours, minutes := fetchers.GetUpcomingBossFromStruct(upcomingBoss)
	bossTime := time.Now().In(loc).Add(time.Minute * time.Duration(upcomingBoss.Time))
	date := bossTime.In(loc).Format(time.TimeOnly)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Next world boss is ***%s*** in %dh%02dm (%s)", name, hours, minutes, date),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
