package mapstore

import (
	"errors"
	"math/rand"
	"tarot-cards-tgbot/internal/constants"
	"tarot-cards-tgbot/internal/model"
	"tarot-cards-tgbot/internal/store"
	"time"
)

type ImageRepository struct {
	store  *Store
	images map[string]*model.Image
}

func (r *ImageRepository) AddImage(i *model.Image) error {
	_, ok := r.images[i.Name]
	if ok {
		return store.ErrImageNameDuplicated
	}

	r.images[i.Name] = i
	return nil
}

func (r *ImageRepository) Find(name string) (*model.Image, bool) {
	image, ok := r.images[name]
	return image, ok
}

func (r *ImageRepository) randomImage() (*model.Image, error) {
	if len(r.images) == 0 {
		return nil, errors.New("images is empty")
	}
	names := make([]string, 0, len(r.images))
	for name, _ := range r.images {
		names = append(names, name)
	}
	rand.Seed(time.Now().UnixNano())
	randName := names[rand.Intn(len(names))]

	return r.images[randName], nil
}

func (r *ImageRepository) RandomImages() ([]*model.Image, error) {
	num := constants.NumImagesResponse
	if num < 2 {
		return nil, errors.New("invalid num images")
	} else if num > len(r.images) {
		return nil, errors.New("num images less then images")
	}

	var res []*model.Image
	check := make(map[string]bool)
	for len(res) < num {
		img, err := r.randomImage()
		if err != nil {
			return nil, err
		}

		if check[img.Name] == false {
			res = append(res, img)
		}
	}
	return res, nil
}
