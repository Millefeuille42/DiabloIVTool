package commands

import (
	"github.com/bwmarrin/discordgo"
)

func wtCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["wt"]; !ok {
		interactionSendError(s, i, "No world tier provided", discordgo.MessageFlagsEphemeral)
		return
	}

	roleName := ""
	switch optionMap["wt"].StringValue() {
	case "1":
		roleName = "World Tier 1"
		break
	case "2":
		roleName = "World Tier 2"
		break
	case "3":
		roleName = "World Tier 3"
		break
	case "4":
		roleName = "World Tier 4"
		break
	default:
		interactionSendError(s, i, "Invalid world tier provided", discordgo.MessageFlagsEphemeral)
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
