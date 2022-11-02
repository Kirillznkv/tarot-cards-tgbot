package sqlstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) AddUser(u *model.User) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (id_tg) VALUES ($1) RETURNING id",
		u.IdTgbot,
	).Scan(&u.IdDatabase); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByTgID(id int64) (*model.User, bool) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, id_tg FROM users WHERE id_tg = $1",
		id,
	).Scan(
		&u.IdDatabase,
		&u.IdTgbot,
	); err != nil {
		return nil, false
	}

	return u, true
}
