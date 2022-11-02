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
	//if s.userRepository == nil {
	//	s.userRepository = &UserRepository{
	//		store: s,
	//	}
	//}

	return s.userRepository
}

func (s *Store) Images() store.ImageRepository {
	//if s.imageRepository == nil {
	//	s.imageRepository = &ImageRepository{
	//		store: s,
	//	}
	//}

	return s.imageRepository
}
