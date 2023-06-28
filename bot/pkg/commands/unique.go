package commands

import (
	"bot/pkg/redisCache"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/url"
	"strings"
)

// type UniqueItemData struct {
//	Name           string   `json:"name"`
//	Class          string   `json:"class,omitempty"`
//	Type           string   `json:"type"`
//	SecondaryStats []string `json:"secondary_stats,omitempty"`
//	TerciaryStats  []string `json:"terciary_stats,omitempty"`
//	Effect         string   `json:"effect"`
//	Flavor         string   `json:"flavor"`
//	Filters        []string `json:"filters,omitempty"`
//}

const uniqueItemImageUrl = "https://rerollcdn.com/DIABLO4/Uniques/2/%s.png"

func buildUniqueItemEmbed(item redisCache.UniqueItemData) *discordgo.MessageEmbed {
	imageUrlName := strings.ToLower(strings.ReplaceAll(item.Name, " ", "_"))
	imageUrlName = url.QueryEscape(imageUrlName)

	ret := &discordgo.MessageEmbed{
		Type: "rich",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: fmt.Sprintf(uniqueItemImageUrl, imageUrlName),
		},
		Color:       13739905,
		Title:       item.Name,
		Description: item.Flavor,
		Fields:      make([]*discordgo.MessageEmbedField, 0),
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Type",
		Value:  item.Type,
		Inline: false,
	})

	if item.Class != "" {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Class",
			Value:  item.Class,
			Inline: false,
		})
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Effect",
		Value:  item.Effect,
		Inline: false,
	})

	if len(item.SecondaryStats) > 0 {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Secondary Stats",
			Value:  strings.Join(item.SecondaryStats, "\n"),
			Inline: false,
		})
	}

	if len(item.TerciaryStats) > 0 {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Terciary Stats",
			Value:  strings.Join(item.TerciaryStats, "\n"),
			Inline: false,
		})
	}

	return ret
}

func buildUniqueItemsChoices(items redisCache.UniqueItemsData) []*discordgo.ApplicationCommandOptionChoice {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0)
	for _, item := range items.Items {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  item.Name,
			Value: item.Name,
		})
	}

	return choices
}

func uniqueCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["name"]; !ok {
		interactionSendError(s, i, "No item provided", discordgo.MessageFlagsEphemeral)
		return
	}

	uniqueItems, err := redisCache.GetUniqueItems()
	if i.Type != discordgo.InteractionApplicationCommand {
		if err != nil {
			return
		}
		choices := buildUniqueItemsChoices(uniqueItems)
		handleAutocomplete(s, i, optionMap["name"].StringValue(), choices)
		return
	}

	if err != nil {
		interactionSendError(s, i, "Error fetching unique items", discordgo.MessageFlagsEphemeral)
		return
	}

	selectedItem := redisCache.UniqueItemData{}
	for _, item := range uniqueItems.Items {
		if item.Name == optionMap["name"].StringValue() {
			selectedItem = item
			break
		}
	}
	if selectedItem.Name == "" {
		interactionSendError(s, i, "Item not found", discordgo.MessageFlagsEphemeral)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    "",
			Components: nil,
			Embeds: []*discordgo.MessageEmbed{
				buildUniqueItemEmbed(selectedItem),
			},
			AllowedMentions: nil,
			Choices:         nil,
			CustomID:        "",
			Title:           "",
		},
	})

	if err != nil {
		log.Println(err)
	}
}
