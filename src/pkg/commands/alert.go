package commands

import (
	"diablo_iv_tool/pkg/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func alertCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if _, ok := optionMap["span"]; !ok || optionMap["span"].StringValue() == "" {
		interactionSendError(s, i, "No timespan provided", discordgo.MessageFlagsEphemeral)
		return
	}

	roleName := ""
	switch optionMap["span"].StringValue() {
	case "morning":
		roleName = "Morning"
		break
	case "day":
		roleName = "Day"
		break
	case "afternoon":
		roleName = "Afternoon"
		break
	case "evening":
		roleName = "Evening"
		break
	default:
		interactionSendError(s, i, "Invalid timespan provided", discordgo.MessageFlagsEphemeral)
	}

	if i.Member == nil {
		return
	}

	err := discord.SetRole(roleName, i.GuildID, i.Member.User.ID, s)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error assigning role", 0)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("You have been assigned the %s role", roleName),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
