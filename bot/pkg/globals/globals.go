package globals

import "github.com/bwmarrin/discordgo"

var (
	Bot *discordgo.Session
)

var DiscordCommands = []*discordgo.ApplicationCommand{
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
	{
		Name:        "channel",
		Description: "Set channel for bot to post in",
	},
	{
		Name:        "help",
		Description: "Get list of available options",
	},
	{
		Name:        "alert",
		Description: "Register for world boss alert on specified timespan",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "span",
				Description: "Time span to register to (/help to see available options)",
				Required:    true,
			},
		},
	},
	{
		Name:        "class",
		Description: "Set your class(es) on the server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "class",
				Description: "Class to set (/help to see available options)",
				Required:    true,
			},
		},
	},
	{
		Name:        "wt",
		Description: "Set your world tier on the server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "wt",
				Description: "world tier to set (/help to see available options)",
				Required:    true,
			},
		},
	},
	{
		Name:        "timezone",
		Description: "Set your timezone for the server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "timezone",
				Description: "TZ Identifier (Europe/Paris format)",
				Required:    true,
			},
		},
	},
}
