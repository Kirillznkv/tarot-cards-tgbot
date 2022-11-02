package main

import (
	"database/sql"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/constants"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/server"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store/sqlstore"
	_ "github.com/lib/pq"
)

func main() {
	db, err := newDB(constants.DatabaseURL)
	defer db.Close()
	store := sqlstore.New(db)

	s, err := server.New(store)
	if err != nil {
		panic(err)
	}

	s.Start()
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
