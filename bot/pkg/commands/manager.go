package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var commandMap = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

func PopulateCommandMap() {
	commandMap["boss"] = upcomingBossCommandHandler
	commandMap["helltide"] = upcomingHelltideCommandHandler

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
