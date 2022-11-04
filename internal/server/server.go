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

type ChatConfigWithUser struct {
	ChatID             int64
	SuperGroupUsername string
	UserID             int64
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
		name := update.Message.From.UserName
		config := tgbotapi.ChatConfigWithUser{
			ChatID: constants.ChannelId,
			UserID: int(id),
		}

		a, err := s.bot.GetChatMember(config)
		if err != nil || a.Status == "left" {
			s.bot.Send(s.buildRegistrationMassage(id))
			continue
		}

		_, ok := s.store.Users().FindByTgID(id)
		if !ok {
			msg = s.buildWelcomeMassage(id, name)
		} else {
			msg = s.buildTarotMassage(id)
		}

		if msg != nil {
			s.bot.Send(msg)
		} else {
			s.error(id, errors.New("error send msg"))
		}
		s.logRequest(id, name, update.Message.Text)
	}
}

func (s *Server) buildRegistrationMassage(id int64) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(id, constants.RegistrationMassage)
	msg.ParseMode = "HTML"
	return msg
}

func (s *Server) buildWelcomeMassage(id int64, name string) tgbotapi.Chattable {
	s.store.Users().AddUser(&model.User{IdChat: id, Name: name})
	msg := tgbotapi.NewPhotoUpload(id, constants.WelcomeImage)
	msg.Caption = constants.WelcomeMassage
	msg.ParseMode = "HTML"
	return msg
}

func (s *Server) buildTarotMassage(id int64) tgbotapi.Chattable {
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

func (s *Server) logRequest(id int64, name, msg string) {
	s.store.LogRequest(id, name, msg)
}

func (s *Server) error(id int64, err error) {
	s.store.Error(id, err)
}
