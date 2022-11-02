package sqlstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_AddUser(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)
	u := &model.User{IdTgbot: 1}
	assert.NoError(t, s.Users().AddUser(u))
	assert.Error(t, s.Users().AddUser(u))
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)
	u, ok := s.Users().FindByTgID(-1)
	assert.Nil(t, u)
	assert.Equal(t, ok, false)

	testUser := &model.User{IdTgbot: 1}
	s.Users().AddUser(testUser)
	u, ok = s.Users().FindByTgID(1)
	assert.Equal(t, u, testUser)
	assert.Equal(t, ok, true)
}
