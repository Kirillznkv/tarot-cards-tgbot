package mapstore

import (
	"github.com/stretchr/testify/assert"
	"tarot-cards-tgbot/internal/model"
	"tarot-cards-tgbot/internal/store"
	"testing"
)

func TestImageRepository_AddImage(t *testing.T) {
	s := New()
	img := &model.Image{
		Name:        "home",
		URL:         "http://home.me",
		Description: "Love home!",
	}
	assert.NoError(t, s.Images().AddImage(img))
	assert.EqualError(t, s.Images().AddImage(img), store.ErrImageNameDuplicated.Error())
}

func TestImageRepository_Find(t *testing.T) {
	s := New()
	img, ok := s.Images().Find("home")
	assert.Nil(t, img)
	assert.Equal(t, ok, false)

	testImage := &model.Image{
		Name:        "home",
		URL:         "http://home.me",
		Description: "Love home!",
	}
	s.Images().AddImage(testImage)
	img, ok = s.Images().Find("home")
	assert.Equal(t, img, testImage)
	assert.Equal(t, ok, true)
}
