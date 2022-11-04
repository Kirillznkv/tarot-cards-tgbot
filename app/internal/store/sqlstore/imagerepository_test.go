package sqlstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageRepository_AddImage(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("images")

	s := New(db)
	img := &model.Image{
		Name:         "home",
		URL:          "http://home.me",
		Description1: "Love home!",
		Description2: "Love home!",
	}
	assert.NoError(t, s.Images().AddImage(img))
	assert.Error(t, s.Images().AddImage(img))
}

func TestImageRepository_Find(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("images")

	s := New(db)
	img, ok := s.Images().FindByName("home")
	assert.Nil(t, img)
	assert.Equal(t, ok, false)

	testImage := &model.Image{
		Name:         "home",
		URL:          "http://home.me",
		Description1: "Love home!",
		Description2: "Love home!",
	}
	assert.NoError(t, s.Images().AddImage(testImage))
	img, ok = s.Images().FindByName("home")
	assert.Equal(t, img.Name, testImage.Name)
	assert.Equal(t, ok, true)
}
