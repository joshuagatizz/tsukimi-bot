package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	token   string
	session *discordgo.Session
}

func NewBot(token string) *Bot {
	// create a session
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Failed creating bot")
	}

	return &Bot{token, s}
}

func (b *Bot) Run() {
	// add the event handlers
	b.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// open session
	b.session.Open()
	defer b.session.Close() // close session, after function termination

	// keep bot running untill there is NO os interruption (ctrl + C)
	log.Print("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func (b *Bot) RegisterCommands() {
	// to register the slash commands
	b.session.Open()
	defer b.session.Close()

	for _, v := range commands {
		_, err := b.session.ApplicationCommandCreate(b.session.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	log.Println("Registering commands succeeded")
}

func (b *Bot) RemoveCommands() {
	// to remove the slash commands
	b.session.Open()
	defer b.session.Close()

	appId := b.session.State.User.ID
	appCommands, err := b.session.ApplicationCommands(appId, "")
	if err != nil {
		log.Fatal("Cannot get application commands")
	}

	for _, v := range appCommands {
		err := b.session.ApplicationCommandDelete(appId, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	log.Println("Removing commands succeeded")
}
