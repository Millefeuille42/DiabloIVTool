package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var commandMap = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
var componentMap = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

func PopulateCommandMap() {
	commandMap["boss"] = upcomingBossCommandHandler
	commandMap["helltide"] = upcomingHelltideCommandHandler

	commandMap["channel"] = channelCommandHandler
	commandMap["timezone"] = timezoneCommandHandler

	commandMap["unique"] = uniqueCommandHandler

	commandMap["roles"] = rolesCommandHandler

	commandMap["help"] = helpCommandHandler

	componentMap["channel_select"] = channelSelectComponentHandler

	componentMap["class_select"] = classSelectComponentHandler
	componentMap["world_tier_select"] = worldTierSelectComponentHandler
	componentMap["alert_select"] = alertSelectComponentHandler

	componentMap["class_button"] = classButtonComponentHandler
	componentMap["world_tier_button"] = worldTierButtonComponentHandler
	componentMap["alert_button"] = alertButtonComponentHandler
	componentMap["role_remove_button"] = handleRoleRemove
}

func CommandManager(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		log.Printf("Received command: %v", i.ApplicationCommandData().Name)
	}

	if i.Type == discordgo.InteractionMessageComponent {
		log.Printf("Received component: %v", i.MessageComponentData().CustomID)
		id := strings.Split(i.MessageComponentData().CustomID, "-")[0]
		if handler, ok := componentMap[id]; ok {
			handler(s, i)
		}
		return
	}

	if handler, ok := commandMap[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}
