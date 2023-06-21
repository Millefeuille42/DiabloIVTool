package commands

import (
	"diablo_iv_tool/pkg/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func classCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if _, ok := optionMap["class"]; !ok || optionMap["class"].StringValue() == "" {
		interactionSendError(s, i, "No class provided")
		return
	}

	roleName := ""
	switch optionMap["class"].StringValue() {
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
		interactionSendError(s, i, "Invalid class provided")
	}

	if i.Member == nil {
		return
	}

	err := discord.SetRole(roleName, i.GuildID, i.Member.User.ID, s)
	if err != nil {
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
