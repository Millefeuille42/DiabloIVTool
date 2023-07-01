package globals

import "github.com/bwmarrin/discordgo"

var (
	Bot *discordgo.Session
)

var DiscordCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "boss",
		Description: "Get upcoming world boss",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "format",
				Description: "Select output format",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "List",
						Value: "list",
					},
				},
			},
		},
	},
	{
		Name:        "helltide",
		Description: "Get upcoming helltide",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "format",
				Description: "Select output format",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "List",
						Value: "list",
					},
				},
			},
		},
	},
	{
		Name:        "channel",
		Description: "Set channel for bot to post in",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "type",
				Description: "Type of alerts to set (/help to see available options)",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "World Boss",
						Value: "boss",
					},
				},
			},
		},
	},
	{
		Name:        "help",
		Description: "Get list of available options",
	},
	{
		Name:        "roles",
		Description: "Set your diablo related roles on the server",
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
	{
		Name:        "unique",
		Description: "Get information about a unique item",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         "name",
				Description:  "Name of the unique item",
				Required:     true,
				Autocomplete: true,
			},
		},
	},
}
