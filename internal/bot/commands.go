package bot

import "github.com/bwmarrin/discordgo"

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Replies with Pong!",
	},
}
