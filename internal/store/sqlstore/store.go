package sqlstore

import (
	"database/sql"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
)

type Store struct {
	db              *sql.DB
	userRepository  *UserRepository
	imageRepository *ImageRepository
}

func New(db *sql.DB) *Store {
	s := &Store{
		db: db,
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	s.imageRepository = &ImageRepository{
		store: s,
	}

	return s
}

func (s *Store) Users() store.UserRepository {
	return s.userRepository
}

func (s *Store) Images() store.ImageRepository {
	return s.imageRepository
}
