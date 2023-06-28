package commands

import (
	"bot/pkg/models"
	"bot/pkg/redisCache"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingBossesHandler(s *discordgo.Session, i *discordgo.InteractionCreate, guild models.GuildModel, loc *time.Location) {
	bosses, err := redisCache.GetUpcomingBosses()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world bosses", 0)
		return
	}

	message := ""
	for _, boss := range bosses.BossEvents {
		date := boss.Time.In(loc).Format(time.RFC850)
		message = fmt.Sprintf("%s***%s***: %s\n", message, boss.Name, date)
	}

	interactionSendResponse(s, i,
		message,
		0,
	)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func upcomingBossCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss", 0)
		return
	}

	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss (invalid location)", 0)
		return
	}

	if _, ok := optionMap["list"]; ok {
		if optionMap["list"].BoolValue() {
			upcomingBossesHandler(s, i, guild, loc)
			return
		}
	}

	worldBoss, err := redisCache.GetWorldBoss()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world boss", 0)
		return
	}

	zone, err := redisCache.GetWorldBossZone()
	if err != nil {
		if err.Error() != "redis: nil" {
			log.Println(err)
			interactionSendError(s, i, "Error fetching upcoming world boss", 0)
			return
		}
	}

	untilBoss := time.Until(worldBoss.StartTime)
	date := worldBoss.StartTime.In(loc).Format(time.TimeOnly)

	hours := int(untilBoss.Hours())
	minutes := int(untilBoss.Minutes()) - (hours * 60)

	if zone == "no-boss" || zone == "" {
		zone = ""
	} else {
		zone = fmt.Sprintf("at %s ", zone)
	}
	interactionSendResponse(s, i,
		fmt.Sprintf("Next world boss is ***%s*** in %dh%02dm %s(%s)",
			worldBoss.Boss,
			hours,
			minutes,
			zone,
			date),
		0,
	)
}
