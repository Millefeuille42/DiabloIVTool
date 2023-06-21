package commands

import (
	"diablo_iv_tool/pkg/fetchers"
	"diablo_iv_tool/pkg/models"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingBossesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world bosses")
		return
	}

	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world bosses (invalid location)")
		return
	}

	bosses, err := fetchers.GetUpcomingBosses(loc)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world bosses")
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: bosses,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
