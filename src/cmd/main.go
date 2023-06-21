package main

import (
	"diablo_iv_tool/pkg/commands"
	"diablo_iv_tool/pkg/database"
	"diablo_iv_tool/pkg/discord"
	"diablo_iv_tool/pkg/globals"
	"diablo_iv_tool/pkg/models"
	"github.com/bwmarrin/discordgo"
	"log"
)

func setUpBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + globals.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discordBot.AddHandler(discord.BotConnected)

	discordBot.AddHandler(discord.GuildJoined)
	discordBot.AddHandler(discord.GuildLeft)

	discordBot.AddHandler(commands.CommandManager)

	err = discordBot.Open()
	if err != nil {
		log.Fatal(err)
	}

	discord.SetUpCloseHandler(discordBot)
	return discordBot
}

func populateDatabase() error {
	guild := models.GuildModel{}
	err := guild.CreateTable()
	if err != nil {
		return err
	}

	role := models.RoleModel{}
	err = role.CreateTable()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var err error

	globals.SetGlobals()

	database.Database, err = database.NewDatabase(globals.DatabaseDriver, globals.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	err = populateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	globals.Bot = setUpBot()
	commands.PopulateCommandMap()

	routine()
}
