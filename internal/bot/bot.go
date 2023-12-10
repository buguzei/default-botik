package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"taxi-bot/internal/bot/db"
	upd "taxi-bot/internal/bot/updates"
)

func StartBot(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	database, err := db.Open()
	if err != nil {
		log.Fatalln(err)
	}

	defer database.Close()

	msgID := make(map[int64]int)
	stage := make(map[int64]string)

	for update := range updates {
		go upd.WorkWithUpdate(bot, update, database, msgID, stage)
	}
}
