package mapstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
)

type UserRepository struct {
	store *Store
	users map[int64]*model.User
}

func (r *UserRepository) AddUser(u *model.User) error {
	_, ok := r.users[u.IdChat]
	if ok {
		return store.ErrUserIdDuplicated
	}

	r.users[u.IdChat] = u
	return nil
}

func (r *UserRepository) FindByTgID(idUser int64) (*model.User, bool) {
	u, ok := r.users[idUser]
	return u, ok
}
