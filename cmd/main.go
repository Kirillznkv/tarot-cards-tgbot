package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io"
	"tarot-cards-tgbot/internal/constants"
)

type MyRequestFileData struct {
	data string
}

func (d *MyRequestFileData) NeedsUpload() bool {
	return false
}

func (d *MyRequestFileData) UploadData() (string, io.Reader, error) {
	return "", nil, nil
}

func (d *MyRequestFileData) SendData() string {
	return string(d.data)
}

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

	imageLinks := []string{
		"https://gfx.tarot.com/images/site/decks/universal-waite/full_size/0.jpg",
		"https://gfx.tarot.com/images/site/decks/universal-waite/full_size/1.jpg",
		"https://gfx.tarot.com/images/site/decks/universal-waite/full_size/2.jpg",
		"https://gfx.tarot.com/images/site/decks/universal-waite/full_size/3.jpg",
		"https://gfx.tarot.com/images/site/decks/universal-waite/full_size/4.jpg",
		"https://gfx.tarot.com/images/site/decks/universal-waite/full_size/5.jpg",
	}
	images := make([]interface{}, 0, len(imageLinks))
	for _, link := range imageLinks {
		images = append(images, tgbotapi.NewInputMediaPhoto(link))
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMediaGroup(update.Message.Chat.ID, images)
		//msg.Caption = "Test"
		bot.Send(msg)
	}
}
