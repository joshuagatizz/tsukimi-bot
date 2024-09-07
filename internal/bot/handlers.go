package bot

import (
	"fmt"
	"log"
	"strings"
	"tsukimi-web/internal/data"
	"tsukimi-web/internal/utils"

	"github.com/bwmarrin/discordgo"
)

const (
	COMMAND_PING = "ping"
	COMMAND_RPS  = "rps"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	COMMAND_PING: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	},
	COMMAND_RPS: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand && i.ID != "" {
			// user id who started the challenge
			userId := utils.GetUser(i).ID

			// their choice (rock, paper, or scissors)
			userChoice := i.Interaction.ApplicationCommandData().Options[0].Value

			// save the state so that when another user accepts we remember what the first person chose
			data.ActiveRockPaperScissorsGame[userId] = &data.RockPaperScissorsGame{
				FirstUserId:     userId,
				FirstUserChoice: userChoice.(string),
			}

			// respond with an interface with a button for someone else to accept the challenge
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Rock papers scissors challenge from <@" + userId + ">",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "Accept",
									Style:    discordgo.PrimaryButton,
									CustomID: "rps-accept_button_" + userId,
								},
							},
						},
					},
				},
			})
		}

		if i.Type == discordgo.InteractionMessageComponent {
			componentId := i.Interaction.MessageComponentData().CustomID
			if strings.HasPrefix(componentId, "rps-accept_button_") {
				// someone has accepted the challenge, give them the rock paper scissors options.
				// + remove "Accept" button in the previous message to prevent someone else clicking it too
				gameId := strings.Replace(componentId, "rps-accept_button_", "", 1)

				secondUserId := utils.GetUser(i).ID
				content := "<@" + secondUserId + "> has accepted the challenge from " + "<@" + gameId + ">"
				s.ChannelMessageEditComplex(&discordgo.MessageEdit{
					ID:         i.Message.ID,
					Channel:    i.Message.ChannelID,
					Content:    &content,
					Components: &[]discordgo.MessageComponent{},
				})

				val, ok := data.ActiveRockPaperScissorsGame[gameId]
				if !ok {
					log.Println("Something went wrong")
					return
				} else {
					val.SecondUserId = secondUserId
				}

				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "What is your object of choice <@" + secondUserId + ">?",
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.SelectMenu{
										MenuType: discordgo.StringSelectMenu,
										CustomID: "rps-select_choice_" + gameId,
										Options: []discordgo.SelectMenuOption{
											{Label: "Rock 🪨", Value: "rock"},
											{Label: "Paper 📄", Value: "paper"},
											{Label: "Scissors ✂️", Value: "scissors"},
										},
									},
								},
							},
						},
					},
				})
			} else if strings.HasPrefix(componentId, "rps-select_choice_") {
				// the second player has picked their choice, now determine the winner
				gameId := strings.Replace(componentId, "rps-select_choice_", "", 1)

				game, ok := data.ActiveRockPaperScissorsGame[gameId]
				if ok {
					secondUserId := utils.GetUser(i).ID
					secondUserChoice := i.Interaction.MessageComponentData().Values[0]

					if secondUserId != game.SecondUserId {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: "You're not the user playing..",
								Flags:   discordgo.MessageFlagsEphemeral,
							},
						})

						return
					}

					// remote the game from storage
					delete(data.ActiveRockPaperScissorsGame, gameId)

					content := "Nice choice " + utils.GetRandomEmoji()
					s.ChannelMessageEditComplex(&discordgo.MessageEdit{
						ID:         i.Message.ID,
						Channel:    i.Message.ChannelID,
						Content:    &content,
						Components: &[]discordgo.MessageComponent{},
					})

					// determine the winner
					result := ""
					if game.FirstUserChoice == secondUserChoice {
						result = fmt.Sprintf(`<@%s> and <@%s> draw with **%s**...`, game.FirstUserId, secondUserId, game.FirstUserChoice)
					} else if (game.FirstUserChoice == "rock" && secondUserChoice == "scissors") || (game.FirstUserChoice == "scissors" && secondUserChoice == "paper") || (game.FirstUserChoice == "paper" && secondUserChoice == "rock") {
						result = fmt.Sprintf(`<@%s>'s **%s** win against <@%s>'s **%s**!`, game.FirstUserId, game.FirstUserChoice, secondUserId, secondUserChoice)
					} else {
						result = fmt.Sprintf(`<@%s>'s **%s** win against <@%s>'s **%s**!`, secondUserId, secondUserChoice, game.FirstUserId, game.FirstUserChoice)
					}

					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{Content: result},
					})
				}
			}
		}
	},
}
