package main

import (
	admin "DobroBot/admin/telegram"
	"DobroBot/client/telegram"
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

		if admin.IsAdmin(update.Message.From.ID) {
			admin.AdminPanel(update, &msg)
		} else {
			msg := tg.NewMessage(update.Message.Chat.ID, "Приветствие и информация ...")
			bot.Send(msg)
			telegram.ClientPanel(update, &msg)
		}
		//bot.Send(msg)
	}
}
