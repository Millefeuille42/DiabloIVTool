package commands

import (
	"bot/pkg/models"
	"bot/pkg/redisCache"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingHelltidesHandler(s *discordgo.Session, i *discordgo.InteractionCreate, guild models.GuildModel, loc *time.Location) {
	helltides, err := redisCache.GetUpcomingHelltides()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltides", 0)
		return
	}

	message := ""

	for _, helltide := range helltides.HelltideEvents {
		date := helltide.In(loc).Format(time.RFC850)
		message = fmt.Sprintf("%s%s\n", message, date)
	}

	interactionSendResponse(s, i,
		fmt.Sprintf("Upcoming helltides:\n```%s```", message),
		0,
	)
}

func upcomingHelltideCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
			upcomingHelltidesHandler(s, i, guild, loc)
			return
		}
	}

	helltide, err := redisCache.GetHelltide()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming helltide", 0)
		return
	}

	untilHelltide := time.Until(helltide.StartTime)
	date := helltide.StartTime.In(loc).Format(time.TimeOnly)

	minutes := int(untilHelltide.Minutes())

	interactionSendResponse(s, i,
		fmt.Sprintf("Next helltide in %02dm (%s)", minutes, date),
		0,
	)
}
