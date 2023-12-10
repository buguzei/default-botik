package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"taxi-bot/config"
	"taxi-bot/internal/bot"
)

func main() {
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

	Bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	config.AdminID, _ = strconv.ParseInt(os.Getenv("ADMIN_TG_ID"), 0, 64)

	bot.StartBot(Bot)
}
