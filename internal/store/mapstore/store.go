package mapstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
)

type Store struct {
	userRepository  *UserRepository
	imageRepository *ImageRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Users() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[int64]*model.User),
		}
	}

	return s.userRepository
}

func (s *Store) Images() store.ImageRepository {
	if s.imageRepository == nil {
		s.imageRepository = &ImageRepository{
			store:  s,
			images: make(map[string]*model.Image),
		}
	}

	return s.imageRepository
}
