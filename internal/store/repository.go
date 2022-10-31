package store

import "tarot-cards-tgbot/internal/model"

type UserRepository interface {
	AddUser(u *model.User) error
	Find(id int64) (*model.User, bool)
}

type ImageRepository interface {
	AddImage(i *model.Image) error
	Find(name string) (*model.Image, bool)
	RandomImages() ([]*model.Image, error)
}
