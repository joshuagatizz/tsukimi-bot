package main

import (
	"log"
	"os"
	"tsukimi-web/internal/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	bot := bot.NewBot(botToken)

	bot.Run()
}
