package server

import (
	"errors"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/constants"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
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

	s.logger.Info("bot started")
	for update := range updates {
		if update.Message == nil {
			continue
		}
		var msg tgbotapi.Chattable

		id := update.Message.Chat.ID
		_, ok := s.store.Users().FindByTgID(id)
		if !ok {
			msg = s.buildWelcomeMassage(id)
		} else {
			msg = s.buildTarorMassage(id)
		}

		if msg != nil {
			s.bot.Send(msg)
		} else {
			s.error(id, errors.New("error send msg"))
		}
		s.logRequest(id, update.Message.Text)
	}
}

func (s *Server) buildWelcomeMassage(id int64) tgbotapi.Chattable {
	s.store.Users().AddUser(&model.User{IdTgbot: id})
	return tgbotapi.NewMessage(id, constants.WelcomeMassage)
}

func (s *Server) buildTarorMassage(id int64) tgbotapi.Chattable {
	imgs, err := s.store.Images().RandomImages()
	if err != nil {
		return nil //need log
	}

	group := model.ImageGroup{}
	for _, img := range imgs {
		group.AddImage(img.InputPhoto)
	}

	return tgbotapi.NewMediaGroup(id, group.Images)
}

func (s *Server) logRequest(id int64, msg string) {
	s.store.LogRequest(id, msg)
}

func (s *Server) error(id int64, err error) {
	s.store.Error(id, err)
}
