package sqlstore

import (
	"errors"
	"fmt"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/constants"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"math/rand"
	"time"
)

type ImageRepository struct {
	store *Store
}

func (r *ImageRepository) AddImage(i *model.Image) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO images (name, url, description) VALUES ($1, $2, $3) RETURNING id",
		i.Name,
		i.URL,
		i.Description,
	).Scan(&i.Id); err != nil {
		return err
	}

	return nil
}

func (r *ImageRepository) FindByName(name string) (*model.Image, bool) {
	img := &model.Image{}

	if err := r.store.db.QueryRow(
		"SELECT id, name, url, description FROM images WHERE name = $1",
		name,
	).Scan(
		&img.Id,
		&img.Name,
		&img.URL,
		&img.Description,
	); err != nil {
		return nil, false
	}

	img.BeforeUse()
	return img, true
}

func (r *ImageRepository) numImages() (int, error) {
	var numImages int
	if err := r.store.db.QueryRow(
		"SELECT COUNT(1) FROM images",
	).Scan(&numImages); err != nil {
		return 0, err
	}

	return numImages, nil
}

func (r *ImageRepository) randomImage() (*model.Image, error) {
	numImages, err := r.numImages()
	if err != nil {
		return nil, err
	}
	if numImages == 0 {
		return nil, errors.New("images is empty")
	}

	rand.Seed(time.Now().UnixNano())
	idRand := rand.Intn(numImages)

	img := &model.Image{}
	if err := r.store.db.QueryRow(
		"SELECT id, name, url, description FROM images ORDER BY id LIMIT 1 OFFSET $1",
		idRand,
	).Scan(
		&img.Id,
		&img.Name,
		&img.URL,
		&img.Description,
	); err != nil {
		return nil, err
	}

	img.BeforeUse()

	return img, nil
}

func (r *ImageRepository) RandomImages() ([]*model.Image, error) {
	numImages, err := r.numImages()
	if err != nil {
		return nil, err
	}
	num := constants.NumImagesResponse

	if num < 2 || num > 10 {
		return nil, errors.New("invalid num images")
	} else if num > numImages {
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
			check[img.Name] = true
		}
	}

	text := ""
	for _, img := range res {
		text += fmt.Sprintf("%s:\n%s\n\n", img.Name, img.Description)
	}
	res[0].InputPhoto.Caption = text

	return res, nil
}
