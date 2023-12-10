package upd

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WorkWithMsg(bot *tgbotapi.BotAPI, update tgbotapi.Update, database *sql.DB, msgID map[int64]int, stage map[int64]string) {
	//userID := update.Message.From.ID
	//username := update.Message.From.FirstName
	//msgText := update.Message.Text
}
