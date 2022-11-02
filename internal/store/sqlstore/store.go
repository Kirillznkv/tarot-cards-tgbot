package sqlstore

import (
	"database/sql"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
	"time"
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

func (s *Store) Error(id int64, err error) {
	s.db.QueryRow(
		"INSERT INTO errors (id_tg, date, time, error) values ($1, $2, $3, $4)",
		id,
		time.Now().UTC(),
		time.Now(),
		err.Error(),
	)
}

func (s *Store) LogRequest(id int64, msg string) {
	s.db.QueryRow(
		"INSERT INTO info (id_tg, date, time, request) values ($1, $2, $3, $4)",
		id,
		time.Now().UTC(),
		time.Now(),
		msg,
	)
}
