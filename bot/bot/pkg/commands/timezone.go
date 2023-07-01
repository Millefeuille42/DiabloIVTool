package commands

import (
	"bot/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func timezoneCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

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

	interactionSendResponse(s, i,
		fmt.Sprintf("Location set to %s", optionMap["timezone"].StringValue()),
		discordgo.MessageFlagsEphemeral,
	)
}
