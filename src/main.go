package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
	"time"
)

const listApiURL = "https://app.worldstone.io:2083/d4"
const apiURL = "https://api.worldstone.io"

type GuildData struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Channel string `json:"channel"`
}

var (
	bot      *discordgo.Session
	guilds   = make(map[string]GuildData, 0)
	commands = [4]*discordgo.ApplicationCommand{
		{
			Name:        "boss",
			Description: "Get upcoming world boss",
		},
		{
			Name:        "bosses",
			Description: "Get list of upcoming world bosses",
		},
		{
			Name:        "helltide",
			Description: "Get upcoming helltide",
		},
		{
			Name:        "helltides",
			Description: "Get list of upcoming helltides",
		},
	}
)

func registerCommands(s *discordgo.Session, g *discordgo.GuildCreate) {
	for _, command := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, g.ID, command)
		if err != nil {
			log.Println(err)
		}
	}
	log.Printf("Registered commands for guild: %v", g.Name)
}

func deleteCommands(s *discordgo.Session, id string) {
	if id == "" {
		for id = range guilds {
			deleteCommands(s, id)
		}
		return
	}

	appCommands, err := s.ApplicationCommands(s.State.User.ID, id)
	if err != nil {
		log.Print(err)
	} else {
		for _, comm := range appCommands {
			_ = s.ApplicationCommandDelete(s.State.User.ID, id, comm.ID)
		}
	}

	if id != "" {
		log.Printf("Deleted commands for guild: %v", id)
	} else {
		log.Printf("Deleted commands for all guilds")
	}
}

func setUpBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	discordBot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	discordBot.AddHandler(func(s *discordgo.Session, g *discordgo.GuildCreate) {
		log.Printf("Joined guild: %v", g.Name)

		guild := GuildData{
			ID:      g.ID,
			Name:    g.Name,
			Channel: "",
		}
		for _, channel := range g.Channels {
			if channel.Type == discordgo.ChannelTypeGuildText {
				permissions, err := s.UserChannelPermissions(s.State.User.ID, channel.ID)
				if err != nil {
					log.Println(err)
					continue
				}

				if (permissions & discordgo.PermissionSendMessages) != 0 {
					guild.Channel = channel.ID
					break
				}
			}
		}
		guilds[g.ID] = guild

		registerCommands(s, g)
	})

	discordBot.AddHandler(func(s *discordgo.Session, g *discordgo.GuildDelete) {
		log.Printf("Left guild: %v", guilds[g.ID].Name)
		delete(guilds, g.ID)
	})

	discordBot.AddHandler(commandHandler)

	err = discordBot.Open()
	if err != nil {
		log.Fatal(err)
	}

	SetUpCloseHandler(discordBot)
	return discordBot
}

func main() {
	bot = setUpBot()
	populateCommandMap()

	shutAlert := false
	lastTime := 0

	for ; ; time.Sleep(1 * time.Minute) {
		upcomingBoss, err := getUpcomingBossData()
		if err != nil {
			log.Println(err)
			continue
		}

		if lastTime < upcomingBoss.Time {
			shutAlert = false
		}
		lastTime = upcomingBoss.Time

		name, hours, minutes := getUpcomingBossFromStruct(upcomingBoss)
		if name == "Wandering Death" {
			name = "Death"
		}

		log.Printf("Next world boss: %s in %dh%02dm (%d)", upcomingBoss.Name, hours, minutes, upcomingBoss.Time)

		err = bot.UpdateWatchStatus(0, fmt.Sprintf("%s in %dh%02dm", name, hours, minutes))
		if err != nil {
			log.Println(err)
		}

		if !shutAlert && upcomingBoss.Time <= 170 {
			log.Printf("Alerting guilds about %s", upcomingBoss.Name)
			for _, guild := range guilds {
				bossTime := time.Now().Local().Add(time.Minute * time.Duration(upcomingBoss.Time))
				date := strings.Replace(bossTime.Local().Format(time.TimeOnly), " CEST", "", -1)

				_, err = bot.ChannelMessageSend(guild.Channel,
					fmt.Sprintf(":warning: World boss ***%s*** in %02dmin (%s)", upcomingBoss.Name, minutes, date))
				if err != nil {
					log.Println(err)
				}
			}
			shutAlert = true
		}
	}
}
