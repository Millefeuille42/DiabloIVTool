package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func classCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["class"]; !ok {
		interactionSendError(s, i, "No class provided", discordgo.MessageFlagsEphemeral)
		return
	}

	optionClass := strings.ToLower(optionMap["class"].StringValue())
	roleName := ""
	switch optionClass {
	case "barbarian":
		roleName = "Barbarian"
		break
	case "sorcerer":
		roleName = "Sorcerer"
		break
	case "rogue":
		roleName = "Rogue"
		break
	case "druid":
		roleName = "Druid"
		break
	case "necromancer":
		roleName = "Necromancer"
		break
	default:
		interactionSendError(s, i, "Invalid class provided", discordgo.MessageFlagsEphemeral)
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
