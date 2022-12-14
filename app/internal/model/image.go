package model

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type Image struct {
	Id           int
	Name         string
	URL          string
	Description1 string
	Description2 string
	InputPhoto   tgbotapi.InputMediaPhoto
}

func (i *Image) BeforeUse() {
	i.InputPhoto = tgbotapi.NewInputMediaPhoto(i.URL)
}
