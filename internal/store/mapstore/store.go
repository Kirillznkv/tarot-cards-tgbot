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
	s := &Store{}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int64]*model.User),
	}
	s.imageRepository = &ImageRepository{
		store:  s,
		images: make(map[string]*model.Image),
	}

	return s
}

func (s *Store) Users() store.UserRepository {
	return s.userRepository
}

func (s *Store) Images() store.ImageRepository {
	return s.imageRepository
}
