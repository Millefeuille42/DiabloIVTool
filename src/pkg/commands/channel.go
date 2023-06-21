package commands

import (
	"diablo_iv_tool/pkg/models"
	"github.com/bwmarrin/discordgo"
	"log"
)

func channelCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error registering channel")
		return
	}

	guild.Channel = i.ChannelID
	err = guild.UpdateGuild()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error registering channel")
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel registered",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
