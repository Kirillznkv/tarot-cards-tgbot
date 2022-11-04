package mapstore

import (
	"fmt"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
	"log"
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
func (s *Store) Error(id int64, err error) {
	log.Fatal(fmt.Printf("user_id: (%d) error: (%s)", id, err))
}
func (s *Store) LogRequest(id int64, name string, msg string) {
	log.Println(fmt.Printf("user_id: (%d) user_name: (%s) request: (%s)", id, name, msg))
}
