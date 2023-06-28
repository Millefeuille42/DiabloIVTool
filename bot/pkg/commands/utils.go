package commands

import (
	"bot/pkg/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	return optionMap
}

func handleRoleRemove(s *discordgo.Session, i *discordgo.InteractionCreate, roleName string) {
	err := discord.UnsetRole(roleName, i.GuildID, i.Member.User.ID, s)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error de-assigning role", 0)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("You have been de-assigned the %s role", roleName),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func handleRoleAdd(s *discordgo.Session, i *discordgo.InteractionCreate, roleName string) {
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
