package telegram

import tg "github.com/go-telegram-bot-api/telegram-bot-api"

func ClientPanel(update tg.Update, msg *tg.MessageConfig) {
	bot, _ := tg.NewBotAPI("6548886185:AAH_D2kYxX2GIV5PhuDWKPjwBpidWeeBVx4")
	keyboard := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("Регестрация"),
			tg.NewKeyboardButton("Задонатить"),
		),
	)
	msg.ReplyMarkup = keyboard

	messageText := update.Message.Text
	switch messageText {
	case "Регестрация":
		// Пользователь нажал кнопку "Добавить промокод"
		// Выполняем действие для кнопки "Добавить промокод"
		// Например, отправляем пользователю сообщение с запросом на ввод промокода
		//msg := tg.NewMessage(update.Message.Chat.ID, "Введите промокод:")
		//bot.Send(msg)
		//promokod := update.Message.Text
		/////TODO add SQL
		//msg = tg.NewMessage(update.Message.Chat.ID, promokod)
		bot.Send(msg)
		///
	case "Задонатить":
		// Пользователь нажал кнопку "Удалить промокод"
		// Выполняем действие для кнопки "Удалить промокод"
		// Например, отправляем пользователю сообщение с запросом на подтверждение удаления промокода
		///TODO delete sql
		msg := tg.NewMessage(update.Message.Chat.ID, ".......\t//https://www.tinkoff.ru/cf/93PoEEuzNTj")
		bot.Send(msg)
	}
}
