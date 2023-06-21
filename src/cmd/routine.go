package main

import (
	"diablo_iv_tool/pkg/fetchers"
	"diablo_iv_tool/pkg/globals"
	"diablo_iv_tool/pkg/models"
	"diablo_iv_tool/pkg/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func setBotStatus(upcomingBoss fetchers.UpcomingBossData) {
	name, hours, minutes := fetchers.GetUpcomingBossFromStruct(upcomingBoss)
	if name == "Wandering Death" {
		name = "Death"
	}

	err := globals.Bot.UpdateWatchStatus(0, fmt.Sprintf("%s in %dh%02dm", name, hours, minutes))
	if err != nil {
		log.Println(err)
	}
	log.Printf("Next world boss: %s in %dh%02dm (%d)", name, hours, minutes, upcomingBoss.Time)
}

func getSpan(upcomingBoss fetchers.UpcomingBossData) string {
	morningStart := utils.GenerateTime(4, 30, time.UTC)
	morningEnd := utils.GenerateTime(6, 30, time.UTC)
	dayStart := utils.GenerateTime(10, 30, time.UTC)
	dayEnd := utils.GenerateTime(12, 30, time.UTC)
	afternoonStart := utils.GenerateTime(16, 30, time.UTC)
	afternoonEnd := utils.GenerateTime(18, 30, time.UTC)
	eveningStart := utils.GenerateTime(22, 30, time.UTC)
	eveningEnd := utils.GenerateTime(0, 30, time.UTC)
	eveningEnd = eveningEnd.AddDate(0, 0, 1)
	now := time.Now().UTC().Add(time.Minute * time.Duration(upcomingBoss.Time))

	switch {
	case now.After(morningStart) && now.Before(morningEnd):
		return "Morning"
	case now.After(dayStart) && now.Before(dayEnd):
		return "Day"
	case now.After(afternoonStart) && now.Before(afternoonEnd):
		return "Afternoon"
	case now.After(eveningStart) && now.Before(eveningEnd):
		return "Evening"
	}

	return ""
}

func getDate(minutes int, location string) (string, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return "", err
	}

	bossTime := time.Now().In(loc).Add(time.Minute * time.Duration(minutes))
	date := bossTime.In(loc).Format(time.TimeOnly)
	return date, nil
}

func getGuildRole(guildId, span string) string {
	roles, err := models.GetRolesByGuildId(guildId)
	if err != nil {
		log.Println(err)
		return ""
	}
	for _, role := range roles {
		if role.Name == span {
			return role.RoleId
		}
	}
	return ""
}

func sendAlert(upcomingBoss fetchers.UpcomingBossData, s *discordgo.Session) {
	span := getSpan(upcomingBoss)
	if span == "" {
		return
	}
	log.Printf("Alerting guilds about %s for %s", upcomingBoss.Name, span)

	guilds, err := models.GetGuilds()
	if err != nil {
		log.Println(err)
		return
	}

	for _, guild := range guilds {
		date, err := getDate(upcomingBoss.Time, guild.Location)
		if err != nil {
			log.Println(err)
			continue
		}

		message := ":warning:"
		selectedRoleId := getGuildRole(guild.GuildId, span)
		if selectedRoleId != "" {
			message += fmt.Sprintf(" <@&%s>", selectedRoleId)
		}
		message += fmt.Sprintf(" World boss ***%s*** in %02dmin (%s)", upcomingBoss.Name, upcomingBoss.Time, date)

		_, err = s.ChannelMessageSend(guild.Channel, message)
		if err != nil {
			log.Println(err)
		}
	}
}

func routine() {
	shutAlert := false
	lastTime := 0

	for ; ; time.Sleep(10 * time.Minute) {
		upcomingBoss, err := fetchers.GetUpcomingBossData()
		if err != nil {
			log.Println(err)
			continue
		}

		if lastTime < upcomingBoss.Time {
			shutAlert = false
		}
		lastTime = upcomingBoss.Time

		setBotStatus(upcomingBoss)

		if !shutAlert && upcomingBoss.Time <= 60 {
			sendAlert(upcomingBoss, globals.Bot)
			shutAlert = true
		}
	}
}
