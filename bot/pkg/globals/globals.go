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
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Morning",
						Value: "morning",
					},
					{
						Name:  "Day",
						Value: "day",
					},
					{
						Name:  "Afternoon",
						Value: "afternoon",
					},
					{
						Name:  "Evening",
						Value: "evening",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "remove",
				Description: "remove role",
				Required:    false,
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
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Barbarian",
						Value: "barbarian",
					},
					{
						Name:  "Sorcerer",
						Value: "sorcerer",
					},
					{
						Name:  "Rogue",
						Value: "rogue",
					},
					{
						Name:  "Druid",
						Value: "druid",
					},
					{
						Name:  "Necromancer",
						Value: "necromancer",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "remove",
				Description: "remove role",
				Required:    false,
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
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "World Tier 1",
						Value: "1",
					},
					{
						Name:  "World Tier 2",
						Value: "2",
					},
					{
						Name:  "World Tier 3",
						Value: "3",
					},
					{
						Name:  "World Tier 4",
						Value: "4",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "remove",
				Description: "remove role",
				Required:    false,
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
