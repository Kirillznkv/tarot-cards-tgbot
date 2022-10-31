package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io/ioutil"
	"tarot-cards-tgbot/internal/constants"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(constants.Token)
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)

	data, _ := ioutil.ReadFile("data/images/the_fool.jpeg")
	b := tgbotapi.FileBytes{Name: "data/images/the_fool.jpeg", Bytes: data}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, b)
		msg.Caption = "Test"
		bot.Send(msg)
	}
}
