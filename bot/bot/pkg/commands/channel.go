package commands

import (
	"bot/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func channelSelectComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	values := i.MessageComponentData().Values
	if len(values) == 0 {
		interactionSendError(s, i, "No channel selected", discordgo.MessageFlagsEphemeral)
		return
	}

	channelTypeArray := strings.Split(i.MessageComponentData().CustomID, "-")
	if len(channelTypeArray) < 2 {
		interactionSendError(s, i, "Error registering channel", 0)
		return
	}
	channelType := channelTypeArray[1]

	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error registering channel", 0)
		return
	}

	switch channelType {
	case "boss":
		guild.BossChannel = values[0]
		break
	case "helltide":
		guild.HelltideChannel = values[0]
		break
	case "legion":
		guild.LegionChannel = values[0]
		break
	}

	err = guild.UpdateGuild()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error registering channel", 0)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Channel <#%s> registered for %s alerts", values[0], channelType),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}

}

func channelCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["type"]; !ok {
		interactionSendError(s, i, "No alert type provided", discordgo.MessageFlagsEphemeral)
		return
	}

	onePointer := 1

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							MenuType:    discordgo.ChannelSelectMenu,
							CustomID:    "channel_select-" + optionMap["type"].StringValue(),
							Placeholder: "Select a channel",
							MinValues:   &onePointer,
							MaxValues:   1,
							ChannelTypes: []discordgo.ChannelType{
								discordgo.ChannelTypeGuildText,
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}
