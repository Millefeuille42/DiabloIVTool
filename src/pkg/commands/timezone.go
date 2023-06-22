package commands

import (
	"diablo_iv_tool/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func timezoneCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if _, ok := optionMap["timezone"]; !ok || optionMap["timezone"].StringValue() == "" {
		interactionSendError(s, i, "No timezone provided", discordgo.MessageFlagsEphemeral)
		return
	}

	_, err := time.LoadLocation(optionMap["timezone"].StringValue())
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Invalid timezone", discordgo.MessageFlagsEphemeral)
		return
	}

	guild := models.GuildModel{}
	_, err = guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error setting location", 0)
		return
	}

	guild.Location = optionMap["timezone"].StringValue()
	err = guild.UpdateGuild()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error setting location", 0)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Location set to %s", optionMap["timezone"].StringValue()),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
