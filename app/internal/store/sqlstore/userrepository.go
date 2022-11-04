package sqlstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) AddUser(u *model.User) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (id_chat, name) VALUES ($1, $2) RETURNING id",
		u.IdChat,
		u.Name,
	).Scan(&u.IdDatabase); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByTgID(id int64) (*model.User, bool) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, id_chat, name FROM users WHERE id_chat = $1",
		id,
	).Scan(
		&u.IdDatabase,
		&u.IdChat,
		&u.Name,
	); err != nil {
		return nil, false
	}

	return u, true
}
