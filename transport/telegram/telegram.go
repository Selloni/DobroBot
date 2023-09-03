package telegram

import (
	"DobroBot/model"
	"DobroBot/store"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mainMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Благотворительность"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("О фонде"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Достижения"),
	),
)

type Telegram struct {
	s        store.Store
	wannaPay map[int64]struct{}
	ch       chan (model.Discont)
	//registration map[int]model.User
}

func NewTelegram(s store.Store, ch chan (model.Discont)) *Telegram {
	return &Telegram{
		s:  s,
		ch: ch,
	}
}

const textAbout = `	Фонд борьбы с диабетом — некоммерческая организация
помощи людям с сахарным диабетом.
	Мы стремимся сделать так, чтобы каждый человек с сахарным диабетом, вне зависимости от своего местонахождения и жизненной ситуации, имел возможность получить помощь
 Наши контакты: +7 (980) 915-72-22 ( WhatsApp или Telegram)
info@diabet-fond.ru
https://diabet-fond.ru/
ул. Сеченова 5, Казань`

const Hello = `Добро пожаловать в "Фонд Борьбы с диабетом"!
🌟 Приветствуем вас в нашем боте, где добро и скидки идут рука об руку! 🌟
🤝 Мы - команда, которая объединила силы, чтобы помочь бороться с редкими заболеваниями у детей. Ваши пожертвования становятся мостом к надежде и лучшему будущему для маленьких борцов!
🎁 В благодарность за вашу поддержку наших маленьких героев, мы предоставляем вам доступ к эксклюзивным скидкам от наших добрых партнеров. Теперь каждая ваша помощь приносит пользу не только другим людям, но и вам лично!`

func (t *Telegram) Run(token string) {
	t.wannaPay = make(map[int64]struct{}, 10)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	go t.checkForDisconts(bot)
	//t.registration = make(map[int]model.User, 10)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		update.Message.IsCommand()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if _, ok := t.wannaPay[update.Message.From.ID]; ok {
			delete(t.wannaPay, update.Message.From.ID)

			count, err := strconv.Atoi(update.Message.Text)
			if err != nil {
				log.Printf("cant parse %v: %v", update.Message.Text, err)
				continue
			}

			// какой то сервис оплаты надо чтоли
			t.s.UpdateDontes(int(update.Message.From.ID), count)
			msg.Text = "https://www.tinkoff.ru/cf/93PoEEuzNTj\n\nСпасибо за пожертвование)"
		} else {
			user, err := t.s.Get(int(update.Message.From.ID))
			if err != nil {
				user = model.User{
					Id:       int(update.Message.From.ID),
					Username: update.Message.From.UserName,
				}
				t.s.Add(user)
			}
			switch update.Message.Text {
			case "/start":
				msg.Text = Hello
				msg.ReplyMarkup = mainMenuKeyboard

			case "Благотворительность":
				msg.Text = "Сколько вы готовы пожертвовать?"
				t.wannaPay[update.Message.From.ID] = struct{}{}
			case "О фонде":
				msg.Text = textAbout
			case "Достижения":
				sumDonate, err2 := t.s.Get(int(update.Message.From.ID))
				if err2 != nil {
					log.Fatalf("couldn't get the donation amount, %v", err)
				}

				msg.Text = fmt.Sprintf("Твой уровень: %s\n Твоя популярность: %d ед",
					t.heroLvl(sumDonate), sumDonate.Donations)
			default:
				msg.Text = "Пожалуйста, используйте команды для взаимодействия с ботом."
			}
		}

		if _, err := bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}
}

func (t *Telegram) heroLvl(donate model.User) string {
	title := "Иследователь"
	if donate.Donations >= 5000 {
		title = "Герой Империи"
	} else if donate.Donations > 1000 {
		title = "Герой Города"
	} else if donate.Donations >= 500 {
		title = "Герой Квартала"
	} else if donate.Donations > 100 {
		title = "Искатель приключений"
	}
	return title
}

func (t *Telegram) checkForDisconts(bot *tgbotapi.BotAPI) {
	for val := range t.ch {
		result, _ := t.s.GetAllWithDonates(val.ForDonate)

		for _, v := range result {
			msg := tgbotapi.NewMessage(int64(v), val.Text)
			msg.ParseMode = "HTML"

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
