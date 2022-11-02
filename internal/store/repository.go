package store

import "github.com/Kirillznkv/tarot-cards-tgbot/internal/model"

type UserRepository interface {
	AddUser(u *model.User) error
	FindByTgID(id int64) (*model.User, bool)
}

type ImageRepository interface {
	AddImage(i *model.Image) error
	FindByName(name string) (*model.Image, bool)
	RandomImages() ([]*model.Image, error)
}
