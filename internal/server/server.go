package server

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"tarot-cards-tgbot/internal/constants"
	"tarot-cards-tgbot/internal/model"
	"tarot-cards-tgbot/internal/store"
)

type Server struct {
	bot    *tgbotapi.BotAPI
	logger *logrus.Logger
	store  store.Store
}

func New(store store.Store) (*Server, error) {
	s := &Server{
		logger: logrus.New(),
		store:  store,
	}

	bot, err := tgbotapi.NewBotAPI(constants.Token)
	if err != nil {
		return nil, err
	}
	s.bot = bot

	return s, nil
}

func (s *Server) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := s.bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		id := update.Message.Chat.ID
		_, ok := s.store.Users().Find(id)
		if !ok {
			msg := s.buildWelcomeMassage(id)
			s.bot.Send(msg)
		} else {
			img, msg := s.buildImageMassage(id)
			if img != nil && msg != nil {
				s.bot.Send(img)
				s.bot.Send(msg)
			}
		}
	}
}

func (s *Server) buildWelcomeMassage(id int64) tgbotapi.Chattable {
	s.store.Users().AddUser(&model.User{ID: id})
	return tgbotapi.NewMessage(id, constants.WelcomeMassage)
}

func (s *Server) buildImageMassage(id int64) (tgbotapi.Chattable, tgbotapi.Chattable) {
	imgs, err := s.store.Images().RandomImages()
	if err != nil {
		return nil, nil //need log
	}

	group := model.ImageGroup{}
	for _, img := range imgs {
		group.AddImage(img.InputPhoto)
		group.Text += fmt.Sprintf("%s:\n%s\n\n", img.Name, img.Description)
	}

	return tgbotapi.NewMediaGroup(id, group.Images), tgbotapi.NewMessage(id, group.Text)
}
