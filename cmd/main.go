package main

import (
	"tarot-cards-tgbot/internal/initstore"
	"tarot-cards-tgbot/internal/server"
	"tarot-cards-tgbot/internal/store/mapstore"
)

func main() {
	store := mapstore.New()
	initstore.Images(store)

	s, err := server.New(store)
	if err != nil {
		panic(err)
	}

	s.Start()
}
