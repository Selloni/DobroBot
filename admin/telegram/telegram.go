package telegram

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

var AdminId = [1]int{789640078}

func IsAdmin(UserId int) bool {
	fmt.Println(UserId)
	for _, id := range AdminId {
		if UserId == id {
			return true
		}
	}
	return false
}

func AdminPanel(update tg.Update) {
	bot, _ := tg.NewBotAPI("6548886185:AAH_D2kYxX2GIV5PhuDWKPjwBpidWeeBVx4")
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("Добавить промокод"),
		tg.NewKeyboardButton("Удалить промокод"),
	)
	//messageText := update.Message.Text
	//switch messageText {
	//case "Добавить промокод":
	//	// Пользователь нажал кнопку "Добавить промокод"
	//	// Выполняем действие для кнопки "Добавить промокод"
	//	// Например, отправляем пользователю сообщение с запросом на ввод промокода
	//	msg := tg.NewMessage(update.Message.Chat.ID, "Введите промокод:")
	//	bot.Send(msg)
	//case "Удалить промокод":
	//	// Пользователь нажал кнопку "Удалить промокод"
	//	// Выполняем действие для кнопки "Удалить промокод"
	//	// Например, отправляем пользователю сообщение с запросом на подтверждение удаления промокода
	//	msg := tg.NewMessage(update.Message.Chat.ID, "Вы уверены, что хотите удалить промокод?")
	//	bot.Send(msg)
	//}

}
