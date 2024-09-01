package main

import (
	"log"
	"os"
	bot "tsukimi-web/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run() // call the run function of bot/bot.go
}
