package main

import (
	admin "DobroBot/admin/telegram"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tg.NewBotAPI("6548886185:AAH_D2kYxX2GIV5PhuDWKPjwBpidWeeBVx4")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		// Получаем Telegram ID пользователя
		fmt.Println(update.Message.From.ID, bot)
		//keyboard := tg.NewReplyKeyboard(
		//	tg.NewKeyboardButtonRow(
		//		tg.NewKeyboardButton("Кнопка 1"),
		//		tg.NewKeyboardButton("Кнопка 2"),
		//	),
		//	tg.NewKeyboardButtonRow(
		//		tg.NewKeyboardButton("Кнопка 3"),
		//		tg.NewKeyboardButton("Кнопка 4"),
		//	),
		//)
		//
		//// Устанавливаем клавиатуру в сообщение
		//msg.ReplyMarkup = keyboard
		if admin.IsAdmin(update.Message.From.ID) {
			admin.AdminPanel(update, &msg)
		} else {

		}
		bot.Send(msg)
	}
}
