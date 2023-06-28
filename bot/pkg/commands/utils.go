package commands

import (
	"bot/pkg/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func interactionSendResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   flags,
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func interactionSendError(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
	interactionSendResponse(s, i, message, flags)
}

func parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	return optionMap
}

func handleRoleRemove(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := discord.UnsetAllRoles(i.GuildID, i.Member.User.ID, s)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error de-assigning roles", 0)
		return
	}

	interactionSendResponse(s, i,
		fmt.Sprintf("You have been de-assigned your roles"),
		discordgo.MessageFlagsEphemeral,
	)
}

func handleRoleAdd(s *discordgo.Session, i *discordgo.InteractionCreate, roleName string) {
	err := discord.SetRole(roleName, i.GuildID, i.Member.User.ID, s)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error assigning role", 0)
		return
	}
}

func handleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate, val string, choices []*discordgo.ApplicationCommandOptionChoice) {
	if i.Type != discordgo.InteractionApplicationCommandAutocomplete {
		return
	}

	choices = filterChoices(choices, val)
	choices = rankChoices(choices, val)
	maxResults := 7
	if len(choices) < maxResults {
		maxResults = len(choices)
	}
	choices = choices[:maxResults]

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
