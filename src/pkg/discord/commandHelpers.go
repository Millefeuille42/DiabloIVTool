package discord

import (
	"diablo_iv_tool/pkg/globals"
	"diablo_iv_tool/pkg/models"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func RegisterCommands(s *discordgo.Session, g *discordgo.GuildCreate) {
	for _, command := range globals.DiscordCommands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, g.ID, command)
		log.Printf("Registered command: %s for guild: %s", command.Name, g.Name)
		time.Sleep(300 * time.Millisecond)
		if err != nil {
			log.Println(err)
		}
	}
	log.Printf("Created discord commands for guild: %v", g.ID)
}

func DeleteCommands(s *discordgo.Session, id string) {
	if id == "" {
		guilds, err := models.GetGuilds()
		if err != nil {
			log.Println(err)
			return
		}
		for _, guild := range guilds {
			DeleteCommands(s, guild.GuildId)
		}
		return
	}

	appCommands, err := s.ApplicationCommands(s.State.User.ID, id)
	if err != nil {
		log.Print(err)
	} else {
		for _, comm := range appCommands {
			_ = s.ApplicationCommandDelete(s.State.User.ID, id, comm.ID)
			log.Printf("Deleted command: %s for guild: %s", comm.Name, id)
		}
	}

	if id != "" {
		log.Printf("Deleted discord commands for guild: %v", id)
	} else {
		log.Printf("Deleted discord commands for all guilds")
	}
}
