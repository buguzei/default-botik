package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewMsg(bot *tgbotapi.BotAPI, chatID int64, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "html"

	msgConfig, err := bot.Send(msg)
	if err != nil {
		return tgbotapi.Message{}, err
	}
	return msgConfig, nil
}

func NewMsgAndMarkup(bot *tgbotapi.BotAPI, chatID int64, text string, kb tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb
	msg.ParseMode = "html"

	msgConfig, err := bot.Send(msg)
	if err != nil {
		return msgConfig, err
	}

	return msgConfig, nil
}

func NewEditMsgMediaCfg(bot *tgbotapi.BotAPI, chatID int64, msgID int, media string) error {
	editMedia := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:    chatID,
			MessageID: msgID,
		},
		Media: tgbotapi.FileID(media),
	}

	_, err := bot.Send(editMedia)
	if err != nil {
		return err
	}

	return nil
}

func NewEditMsgAndMarkup(bot *tgbotapi.BotAPI, chatID int64, msgID int, text string, kb tgbotapi.InlineKeyboardMarkup, needKb bool) error {
	editedMsg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, text, kb)
	editedMsg.ParseMode = "html"

	if !needKb {
		editedMsg.ReplyMarkup = nil
	}

	_, err := bot.Send(editedMsg)
	if err != nil {
		return err
	}

	return nil
}

func NewDelMsg(bot *tgbotapi.BotAPI, chatID int64, msgID int) {
	delMsg := tgbotapi.NewDeleteMessage(chatID, msgID)

	bot.Send(delMsg)
}

func NewPhoto(bot *tgbotapi.BotAPI, chatID int64, file tgbotapi.RequestFileData, caption string) error {
	newPh := tgbotapi.NewPhoto(chatID, file)
	newPh.Caption = caption

	_, err := bot.Send(newPh)
	if err != nil {
		return err
	}

	return nil
}

func NewPhotoAndMarkup(bot *tgbotapi.BotAPI, chatID int64, file tgbotapi.RequestFileData, caption string, kb tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	newPh := tgbotapi.NewPhoto(chatID, file)
	newPh.Caption = caption
	newPh.ReplyMarkup = kb

	msgConfig, err := bot.Send(newPh)
	if err != nil {
		return tgbotapi.Message{}, err
	}

	return msgConfig, nil
}

func NewAlert(bot *tgbotapi.BotAPI, id, text string) error {
	alertMsg := tgbotapi.NewCallbackWithAlert(id, text)

	_, err := bot.Send(alertMsg)
	if err != nil {
		return err
	}

	return nil
}
