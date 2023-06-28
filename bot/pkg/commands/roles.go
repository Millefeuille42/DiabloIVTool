package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func rolesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Choose a class",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Classes",
							Style: discordgo.PrimaryButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🎲",
							},
							CustomID: "class_button",
						},
						discordgo.Button{
							Label: "World tiers",
							Style: discordgo.PrimaryButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🌎",
							},
							CustomID: "world_tier_button",
						},
						discordgo.Button{
							Label: "Alerts",
							Style: discordgo.PrimaryButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🚨",
							},
							CustomID: "alert_button",
						},
						discordgo.Button{
							Label: "Remove",
							Style: discordgo.DangerButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "❌",
							},
							CustomID: "role_remove_button",
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}

func classSelectComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	values := i.MessageComponentData().Values

	if len(values) == 0 {
		interactionSendError(s, i, "No class selected", discordgo.MessageFlagsEphemeral)
		return
	}

	if i.Member == nil {
		return
	}

	roleName := ""
	for _, value := range values {
		switch value {
		case "barbarian":
			roleName = "Barbarian"
			break
		case "sorcerer":
			roleName = "Sorcerer"
			break
		case "rogue":
			roleName = "Rogue"
			break
		case "druid":
			roleName = "Druid"
			break
		case "necromancer":
			roleName = "Necromancer"
			break
		default:
			interactionSendError(s, i, "Invalid class provided", discordgo.MessageFlagsEphemeral)
		}

		handleRoleAdd(s, i, roleName)
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "You have been assigned the selected roles",
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func worldTierSelectComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	values := i.MessageComponentData().Values

	if len(values) == 0 {
		interactionSendError(s, i, "No class selected", discordgo.MessageFlagsEphemeral)
		return
	}

	if i.Member == nil {
		return
	}

	for _, value := range values {
		roleName := ""
		switch value {
		case "1":
			roleName = "World Tier 1"
			break
		case "2":
			roleName = "World Tier 2"
			break
		case "3":
			roleName = "World Tier 3"
			break
		case "4":
			roleName = "World Tier 4"
			break
		default:
			interactionSendError(s, i, "Invalid world tier provided", discordgo.MessageFlagsEphemeral)
		}

		handleRoleAdd(s, i, roleName)
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "You have been assigned the selected roles",
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func alertSelectComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	values := i.MessageComponentData().Values

	if len(values) == 0 {
		interactionSendError(s, i, "No class selected", discordgo.MessageFlagsEphemeral)
		return
	}

	if i.Member == nil {
		return
	}

	for _, value := range values {
		roleName := ""
		switch value {
		case "morning":
			roleName = "Morning"
			break
		case "day":
			roleName = "Day"
			break
		case "afternoon":
			roleName = "Afternoon"
			break
		case "evening":
			roleName = "Evening"
			break
		default:
			interactionSendError(s, i, "Invalid timespan provided", discordgo.MessageFlagsEphemeral)
		}

		handleRoleAdd(s, i, roleName)
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "You have been assigned the selected roles",
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func classButtonComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	onePointer := 1

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							MenuType:    discordgo.StringSelectMenu,
							CustomID:    "class_select",
							Placeholder: "Select your class(es)",
							MinValues:   &onePointer,
							MaxValues:   5,
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Barbarian",
									Value: "barbarian",
									Emoji: discordgo.ComponentEmoji{
										Name: "🗡️",
									},
								},
								{
									Label: "Sorcerer",
									Value: "sorcerer",
									Emoji: discordgo.ComponentEmoji{
										Name: "🧙",
									},
								},
								{
									Label: "Rogue",
									Value: "rogue",
									Emoji: discordgo.ComponentEmoji{
										Name: "🏹",
									},
								},
								{
									Label: "Druid",
									Value: "druid",
									Emoji: discordgo.ComponentEmoji{
										Name: "🌳",
									},
								},
								{
									Label: "Necromancer",
									Value: "necromancer",
									Emoji: discordgo.ComponentEmoji{
										Name: "💀",
									},
								},
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}

func worldTierButtonComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	onePointer := 1

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							MenuType:    discordgo.StringSelectMenu,
							CustomID:    "world_tier_select",
							Placeholder: "Select your world tier(s)",
							MinValues:   &onePointer,
							MaxValues:   4,
							Options: []discordgo.SelectMenuOption{
								{
									Label: "World Tier 1",
									Value: "1",
									Emoji: discordgo.ComponentEmoji{
										Name: "1️⃣",
									},
								},
								{
									Label: "World Tier 2",
									Value: "2",
									Emoji: discordgo.ComponentEmoji{
										Name: "2️⃣",
									},
								},
								{
									Label: "World Tier 3",
									Value: "3",
									Emoji: discordgo.ComponentEmoji{
										Name: "3️⃣",
									},
								},
								{
									Label: "World Tier 4",
									Value: "4",
									Emoji: discordgo.ComponentEmoji{
										Name: "4️⃣",
									},
								},
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}

func alertButtonComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	onePointer := 1

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Select the world boss spawns you want te be alerted for (for details use the /help command)",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							MenuType:    discordgo.StringSelectMenu,
							CustomID:    "alert_select",
							Placeholder: "Select your timespan(s)",
							MinValues:   &onePointer,
							MaxValues:   4,
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Morning",
									Value: "morning",
								},
								{
									Label: "Day",
									Value: "day",
								},
								{
									Label: "Afternoon",
									Value: "afternoon",
								},
								{
									Label: "Evening",
									Value: "evening",
								},
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}
