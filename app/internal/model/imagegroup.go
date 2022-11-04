package model

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type ImageGroup struct {
	Images []interface{}
}

func (g *ImageGroup) AddImage(img tgbotapi.InputMediaPhoto) {
	g.Images = append(g.Images, img)
}
