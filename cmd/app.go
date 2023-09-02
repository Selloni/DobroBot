package main

import (
	"DobroBot/admin/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6548886185:AAH_D2kYxX2GIV5PhuDWKPjwBpidWeeBVx4")
	if err != nil {
		log.Fatal(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		// Получаем Telegram ID пользователя
		if telegram.IsAdmin(update.Message.From.ID) {
			// Создаем новую клавиатуру

		}
		bot.Send(msg)
	}
}
