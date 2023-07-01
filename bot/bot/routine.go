package main

import (
	"bot/pkg/globals"
	"bot/pkg/models"
	"bot/pkg/redisCache"
	"bot/pkg/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func setBotStatus(worldBoss redisCache.WorldBossData, zone string) {
	if worldBoss.Boss == "Wandering Death" {
		worldBoss.Boss = "Death"
	}

	untilBoss := time.Until(worldBoss.StartTime)
	hours := int(untilBoss.Hours())
	minutes := int(untilBoss.Minutes()) - (hours * 60)

	err := globals.Bot.UpdateWatchStatus(0, fmt.Sprintf("%s in %dh%02dm%s",
		worldBoss.Boss,
		hours,
		minutes,
		zone))
	if err != nil {
		log.Println(err)
	}
}

func getSpan(upcomingBoss redisCache.WorldBossData) string {
	morningStart := utils.GenerateTime(4, 30, time.UTC)
	morningEnd := utils.GenerateTime(6, 30, time.UTC)
	dayStart := utils.GenerateTime(10, 30, time.UTC)
	dayEnd := utils.GenerateTime(12, 30, time.UTC)
	afternoonStart := utils.GenerateTime(16, 30, time.UTC)
	afternoonEnd := utils.GenerateTime(18, 30, time.UTC)
	eveningStart := utils.GenerateTime(22, 30, time.UTC)
	eveningEnd := utils.GenerateTime(0, 30, time.UTC)
	eveningEnd = eveningEnd.AddDate(0, 0, 1)

	switch {
	case upcomingBoss.StartTime.After(morningStart) && upcomingBoss.StartTime.Before(morningEnd):
		return "Morning"
	case upcomingBoss.StartTime.After(dayStart) && upcomingBoss.StartTime.Before(dayEnd):
		return "Day"
	case upcomingBoss.StartTime.After(afternoonStart) && upcomingBoss.StartTime.Before(afternoonEnd):
		return "Afternoon"
	case upcomingBoss.StartTime.After(eveningStart) && upcomingBoss.StartTime.Before(eveningEnd):
		return "Evening"
	}

	return ""
}

func getDate(bossTime time.Time, location string) (string, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return "", err
	}

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

func sendAlert(worldBoss redisCache.WorldBossData, s *discordgo.Session, emote string, zone string) {
	span := getSpan(worldBoss)
	if span == "" {
		return
	}
	log.Printf("Alerting guilds about %s for %s", worldBoss.Boss, span)

	guilds, err := models.GetGuilds()
	if err != nil {
		log.Println(err)
		return
	}

	minutes := int(time.Until(worldBoss.StartTime).Minutes())

	for _, guild := range guilds {
		date, err := getDate(worldBoss.StartTime, guild.Location)

		message := emote
		selectedRoleId := getGuildRole(guild.GuildId, span)
		if selectedRoleId != "" {
			message += fmt.Sprintf(" <@&%s>", selectedRoleId)
		}
		message += fmt.Sprintf(" World boss ***%s*** in %02dmin%s (%s)",
			worldBoss.Boss,
			minutes,
			zone,
			date)

		_, err = s.ChannelMessageSend(guild.BossChannel, message)
		if err != nil {
			log.Println(err)
		}
	}
}

func routineEvents(worldBoss redisCache.WorldBossData, zone string) {
	setBotStatus(worldBoss, zone)

	untilBoss := time.Until(worldBoss.StartTime)

	hourBefore := time.NewTimer(untilBoss - time.Hour - time.Second*1)
	defer hourBefore.Stop()
	tenMinutesBefore := time.NewTimer(untilBoss - time.Minute*10 - time.Second*1)
	defer tenMinutesBefore.Stop()
	refreshTimer := time.NewTimer(untilBoss + time.Minute*10)
	defer refreshTimer.Stop()

	updateBotStatusTicker := time.NewTicker(time.Minute * 1)
	defer updateBotStatusTicker.Stop()

	for refresh := false; !refresh; {
		select {
		case <-updateBotStatusTicker.C:
			setBotStatus(worldBoss, zone)
		case <-hourBefore.C:
			sendAlert(worldBoss, globals.Bot, ":warning:", zone)
		case <-tenMinutesBefore.C:
			sendAlert(worldBoss, globals.Bot, ":sos:", zone)
		case <-refreshTimer.C:
			refresh = true
		}
	}
}

func routine() {
	for {
		worldBoss, err := redisCache.GetWorldBoss()
		if err != nil {
			if err.Error() != "redis: nil" {
				log.Println(err)
			}
			continue
		}
		zone, err := redisCache.GetWorldBossZone()
		if err != nil {
			if err.Error() != "redis: nil" {
				log.Println(err)
			}
		}

		if zone == "no-boss" || zone == "" {
			zone = ""
		} else {
			zone = fmt.Sprintf(" at %s", zone)
		}

		routineEvents(worldBoss, zone)
	}
}
