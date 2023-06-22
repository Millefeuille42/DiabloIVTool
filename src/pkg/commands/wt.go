package commands

import (
	"diablo_iv_tool/pkg/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func wtCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if _, ok := optionMap["wt"]; !ok || optionMap["wt"].StringValue() == "" {
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
