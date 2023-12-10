package upd

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Order = make(map[int64]map[string]string)

func WorkWithUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, database *sql.DB, msgID map[int64]int, stage map[int64]string) {

	if update.CallbackQuery != nil {
		WorkWithCallback(bot, update, database, msgID, stage)
	}

	if update.Message != nil {
		if update.Message.Chat.IsPrivate() {
			WorkWithMsg(bot, update, database, msgID, stage)
		}
	}
}
