package model

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type Image struct {
	Name        string
	URL         string
	Description string
	InputPhoto  tgbotapi.InputMediaPhoto
}

func (i *Image) BeforeUse() {
	i.InputPhoto = tgbotapi.NewInputMediaPhoto(i.URL)
}
