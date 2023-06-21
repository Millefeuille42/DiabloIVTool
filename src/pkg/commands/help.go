package commands

import (
	"diablo_iv_tool/pkg/models"
	"diablo_iv_tool/pkg/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func generateSpansString(guild models.GuildModel) (string, error) {
	loc, err := time.LoadLocation(guild.Location)
	if err != nil {
		return "", err
	}

	morningStart := utils.GenerateTime(4, 30, loc)
	morningEnd := utils.GenerateTime(6, 30, loc)
	dayStart := utils.GenerateTime(10, 30, loc)
	dayEnd := utils.GenerateTime(12, 30, loc)
	afternoonStart := utils.GenerateTime(16, 30, loc)
	afternoonEnd := utils.GenerateTime(18, 30, loc)
	eveningStart := utils.GenerateTime(22, 30, loc)
	eveningEnd := utils.GenerateTime(0, 30, loc)

	message := fmt.Sprintf("World boss spawn timespans:\n")
	message += fmt.Sprintf("morning:   %s - %s\n", morningStart.Format(time.Kitchen), morningEnd.Format(time.Kitchen))
	message += fmt.Sprintf("day:       %s - %s\n", dayStart.Format(time.Kitchen), dayEnd.Format(time.Kitchen))
	message += fmt.Sprintf("afternoon: %s - %s\n", afternoonStart.Format(time.Kitchen), afternoonEnd.Format(time.Kitchen))
	message += fmt.Sprintf("evening:   %s - %s\n", eveningStart.Format(time.Kitchen), eveningEnd.Format(time.Kitchen))

	return message, nil
}

func generateWtsString() string {
	return "World Tiers:\n" +
		"1. adventurer\n" +
		"2. veteran\n" +
		"3. nightmare\n" +
		"4. torment\n"
}

func generateClassesString() string {
	return "Classes:\n" +
		"1. barbarian\n" +
		"2. rogue\n" +
		"3. sorcerer\n" +
		"4. druid\n" +
		"5. necromancer\n"
}

func helpCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error generating help")
		return
	}

	message, err := generateSpansString(guild)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error generating help (wrong timezone)")
		return
	}
	message += "\n" + generateWtsString() + "\n" + generateClassesString()
	message = fmt.Sprintf("```%s```", message)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
