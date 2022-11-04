package mapstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageRepository_AddImage(t *testing.T) {
	s := New()
	img := &model.Image{
		Name:         "home",
		URL:          "http://home.me",
		Description1: "Love home!",
		Description2: "Love home!",
	}
	assert.NoError(t, s.Images().AddImage(img))
	assert.EqualError(t, s.Images().AddImage(img), store.ErrImageNameDuplicated.Error())
}

func TestImageRepository_Find(t *testing.T) {
	s := New()
	img, ok := s.Images().FindByName("home")
	assert.Nil(t, img)
	assert.Equal(t, ok, false)

	testImage := &model.Image{
		Name:         "home",
		URL:          "http://home.me",
		Description1: "Love home!",
		Description2: "Love home!",
	}
	err := s.Images().AddImage(testImage)
	assert.NoError(t, err)
	img, ok = s.Images().FindByName("home")
	assert.Equal(t, img, testImage)
	assert.Equal(t, ok, true)
}
