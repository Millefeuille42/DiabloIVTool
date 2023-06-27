package commands

import (
	"bot/pkg/models"
	"bot/pkg/redisCache"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func upcomingBossesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world bosses", 0)
		return
	}

	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world bosses (invalid location)", 0)
		return
	}

	bosses, err := redisCache.GetUpcomingBosses()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error fetching upcoming world bosses", 0)
		return
	}

	message := ":warning: **This info might might be inaccurate** :warning:\n\n"
	for _, boss := range bosses.BossEvents {
		date := boss.Time.In(loc).Format(time.RFC850)
		message = fmt.Sprintf("%s***%s***: %s\n", message, boss.Name, date)
	}

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
