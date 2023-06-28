package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func alertCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if _, ok := optionMap["span"]; !ok {
		interactionSendError(s, i, "No timespan provided", discordgo.MessageFlagsEphemeral)
		return
	}
	optionSpan := strings.ToLower(optionMap["span"].StringValue())

	roleName := ""
	switch optionSpan {
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

	optionRemove := false
	if _, ok := optionMap["remove"]; ok {
		optionRemove = optionMap["remove"].BoolValue()
	}

	if optionRemove {
		handleRoleRemove(s, i, roleName)
	} else {
		handleRoleAdd(s, i, roleName)
	}
}
