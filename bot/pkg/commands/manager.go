package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var commandMap = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

func interactionSendError(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
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

func PopulateCommandMap() {
	commandMap["boss"] = upcomingBossCommandHandler
	commandMap["bosses"] = upcomingBossesCommandHandler
	commandMap["helltide"] = upcomingHelltideCommandHandler
	commandMap["helltides"] = upcomingHelltidesCommandHandler

	commandMap["channel"] = channelCommandHandler
	commandMap["timezone"] = timezoneCommandHandler

	commandMap["alert"] = alertCommandHandler
	commandMap["class"] = classCommandHandler
	commandMap["wt"] = wtCommandHandler

	commandMap["help"] = helpCommandHandler
}

func CommandManager(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("Received command: %v", i.ApplicationCommandData().Name)
	if handler, ok := commandMap[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}
