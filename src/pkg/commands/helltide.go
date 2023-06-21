package commands

import (
	"diablo_iv_tool/pkg/fetchers"
	"diablo_iv_tool/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingHelltideCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helltide, err := fetchers.GetUpcomingHelltideData()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltide")
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

	hours, minutes := fetchers.GetUpcomingHelltideFromStruct(helltide)
	helltideTime := time.Now().In(loc).Add(time.Minute * time.Duration(helltide.Time))
	date := helltideTime.In(loc).Format(time.TimeOnly)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Next helltide in %dh%02dm (%s)", hours, minutes, date),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
