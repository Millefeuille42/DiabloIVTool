package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

var commandMap = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

func populateCommandMap() {
	commandMap["boss"] = upcomingBossCommandHandler
	commandMap["bosses"] = upcomingBossesCommandHandler
	commandMap["helltide"] = upcomingHelltideCommandHandler
	commandMap["helltides"] = upcomingHelltidesCommandHandler
}

func upcomingBossCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	upcomingBoss, err := getUpcomingBossData()
	if err != nil {
		log.Println(err)
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error getting upcoming world boss",
			},
		})
		if err != nil {
			log.Println(err)
		}
	}

	name, hours, minutes := getUpcomingBossFromStruct(upcomingBoss)
	bossTime := time.Now().Local().Add(time.Minute * time.Duration(upcomingBoss.Time))
	date := strings.Replace(bossTime.Local().Format(time.TimeOnly), " CEST", "", -1)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Next world boss is ***%s*** in %dh%02dm (%s)", name, hours, minutes, date),
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func upcomingBossesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	bosses, err := getUpcomingBosses()
	if err != nil {
		log.Println(err)
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error getting upcoming world bosses",
			},
		})
		if err != nil {
			log.Println(err)
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: bosses,
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func upcomingHelltideCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helltide, err := getUpcomingHelltideData()
	if err != nil {
		log.Println(err)
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error getting upcoming helltide",
			},
		})
		if err != nil {
			log.Println(err)
		}
	}

	hours, minutes := getUpcomingHelltideFromStruct(helltide)
	helltideTime := time.Now().Local().Add(time.Minute * time.Duration(helltide.Time))
	date := strings.Replace(helltideTime.Local().Format(time.TimeOnly), " CEST", "", -1)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Next helltide in %dh%02dm (%s)", hours, minutes, date),
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func upcomingHelltidesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helltides, err := getUpcomingHelltides()
	if err != nil {
		log.Println(err)
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error getting upcoming helltides",
			},
		})
		if err != nil {
			log.Println(err)
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Upcoming helltides:\n```%s```", helltides),
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("Received command: %v", i.ApplicationCommandData().Name)
	if handler, ok := commandMap[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}
