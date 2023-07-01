package main

import (
	"bot/pkg/commands"
	"bot/pkg/database"
	"bot/pkg/discord"
	"bot/pkg/globals"
	"bot/pkg/models"
	"bot/pkg/redisCache"
	"github.com/bwmarrin/discordgo"
	"github.com/redis/go-redis/v9"
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

	redisCache.Context = redisCache.NewContext()

	redisCache.Client = redisCache.New(&redis.Options{
		Addr:         globals.RedisHost + ":" + globals.RedisPort,
		Password:     globals.RedisPassword,
		DB:           globals.RedisDB,
		ClientName:   "dbivtool-bot",
		MaxIdleConns: 5,
	})

	defer redisCache.Client.Close()

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
