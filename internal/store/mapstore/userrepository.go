package mapstore

import (
	"tarot-cards-tgbot/internal/model"
	"tarot-cards-tgbot/internal/store"
)

type UserRepository struct {
	store *Store
	users map[int64]*model.User
}

func (r *UserRepository) AddUser(u *model.User) error {
	_, ok := r.users[u.ID]
	if ok {
		return store.ErrUserIdDuplicated
	}

	r.users[u.ID] = u
	return nil
}

func (r *UserRepository) Find(idUser int64) (*model.User, bool) {
	u, ok := r.users[idUser]
	return u, ok
}
