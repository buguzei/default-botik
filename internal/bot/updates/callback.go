package upd

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var OrderMsg = make(map[int64]string)

func WorkWithCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, database *sql.DB, msgID map[int64]int, stage map[int64]string) {
	//userID := update.CallbackQuery.From.ID
	//username := update.CallbackQuery.From.FirstName
}
