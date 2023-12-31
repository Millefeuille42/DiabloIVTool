package discord

import (
	"bot/pkg/models"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func BotConnected(s *discordgo.Session, r *discordgo.Ready) {
	r = nil
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

func GuildJoined(s *discordgo.Session, g *discordgo.GuildCreate) {
	log.Printf("Joined guild: %v", g.Name)

	guild := models.GuildModel{
		Name:    g.Name,
		GuildId: g.ID,
	}
	for _, channel := range g.Channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			permissions, err := s.UserChannelPermissions(s.State.User.ID, channel.ID)
			if err != nil {
				log.Println(err)
				continue
			}

			if (permissions & discordgo.PermissionSendMessages) != 0 {
				guild.BossChannel = channel.ID
				guild.HelltideChannel = channel.ID
				guild.LegionChannel = channel.ID
				break
			}
		}
	}

	guildRegistered := false
	err := guild.CreateGuild()
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println(err)
		} else {
			guildRegistered = true
		}
	}

	RegisterCommands(s, g)
	if !guildRegistered {
		err = CreateRoles(s, g.ID)
		if err != nil {
			log.Println(err)
		}
		return
	}
}

func GuildLeft(s *discordgo.Session, g *discordgo.GuildDelete) {
	log.Printf("Left guild: %v", g.BeforeDelete.Name)
	DeleteCommands(s, g.ID)
	_ = DeleteRoles(s, g.ID)
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(g.ID)
	if err != nil {
		log.Println(err)
		return
	}
	err = guild.DeleteGuild()
	if err != nil {
		log.Println(err)
		return
	}
}
