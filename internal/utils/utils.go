package utils

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func Shuffle[T any](array []T) []T {
	n := len(array)
	for i := range array {
		j := rand.Intn(n)
		array[i], array[j] = array[j], array[i]
	}

	return array
}

func GetUser(i *discordgo.InteractionCreate) *discordgo.User {
	if i.Interaction.User != nil {
		return i.Interaction.User
	}

	return i.Interaction.Member.User
}

func GetRandomEmoji() string {
	emojis := []rune{'😭', '😄', '😌', '🤓', '😎', '😤', '🤖', '🌏', '📸', '💿', '👋', '🌊', '✨'}

	return string(emojis[rand.Intn(len(emojis))])
}
