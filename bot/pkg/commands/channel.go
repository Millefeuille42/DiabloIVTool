package commands

import (
	"bot/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func channelCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["type"]; !ok {
		interactionSendError(s, i, "No alert type provided", discordgo.MessageFlagsEphemeral)
		return
	}

	if _, ok := optionMap["channel"]; !ok {
		interactionSendError(s, i, "No channel provided", discordgo.MessageFlagsEphemeral)
		return
	}

	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error registering channel", 0)
		return
	}

	channelOption := optionMap["channel"].ChannelValue(s)
	if channelOption.Type != discordgo.ChannelTypeGuildText {
		interactionSendError(s, i, "Invalid channel provided", discordgo.MessageFlagsEphemeral)
		return
	}

	typeOption := strings.ToLower(optionMap["type"].StringValue())
	switch typeOption {
	case "boss":
		guild.BossChannel = channelOption.ID
		break
	case "helltide":
		guild.HelltideChannel = channelOption.ID
		break
	case "legion":
		guild.LegionChannel = channelOption.ID
		break
	}

	err = guild.UpdateGuild()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error registering channel", 0)
		return
	}

	interactionSendResponse(s, i,
		fmt.Sprintf("Channel <#%s> registered for %s alerts", channelOption.ID, typeOption),
		discordgo.MessageFlagsEphemeral,
	)
}
