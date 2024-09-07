package bot

import (
	"tsukimi-web/internal/utils"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        COMMAND_PING,
		Description: "Replies with Pong!",
	},
	{
		Name:        COMMAND_RPS,
		Description: "Start a rock paper scissors match.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "String option",
				Required:    true,
				Choices: utils.Shuffle([]*discordgo.ApplicationCommandOptionChoice{
					{Name: "Rock 🪨", Value: "rock"},
					{Name: "Paper 📄", Value: "paper"},
					{Name: "Scissors ✂️", Value: "scissors"},
				}),
			},
		},
	},
}
