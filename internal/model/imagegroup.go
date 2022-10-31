package model

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type ImageGroup struct {
	Text   string
	Images []interface{}
}

func (g *ImageGroup) AddImage(img tgbotapi.InputMediaPhoto) {
	g.Images = append(g.Images, img)
}
