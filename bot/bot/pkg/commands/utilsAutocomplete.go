package commands

import (
	"github.com/bwmarrin/discordgo"
	"sort"
	"strings"
)

func filterChoices(choices []*discordgo.ApplicationCommandOptionChoice, userInput string) []*discordgo.ApplicationCommandOptionChoice {
	var matchingOptions []*discordgo.ApplicationCommandOptionChoice

	for _, choice := range choices {
		if strings.Contains(strings.ToLower(choice.Name), strings.ToLower(userInput)) {
			matchingOptions = append(matchingOptions, choice)
		}
	}

	return matchingOptions
}

func rankChoices(choices []*discordgo.ApplicationCommandOptionChoice, userInput string) []*discordgo.ApplicationCommandOptionChoice {
	scores := make(map[string]int)
	for _, choice := range choices {
		score := calculateSimilarityScore(choice, userInput)
		scores[choice.Name] = score
	}

	sort.Slice(choices, func(i, j int) bool {
		return scores[choices[i].Name] > scores[choices[j].Name]
	})

	return choices
}

func calculateSimilarityScore(choice *discordgo.ApplicationCommandOptionChoice, userInput string) int {
	score := 0

	userInput = strings.ToLower(userInput)
	for _, char := range userInput {
		name := strings.ToLower(choice.Name)
		if strings.ContainsRune(name, char) {
			score++
		}
	}

	return score
}
