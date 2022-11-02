package mapstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_AddUser(t *testing.T) {
	s := New()
	u := &model.User{IdTgbot: 1}
	assert.NoError(t, s.Users().AddUser(u))
	assert.EqualError(t, s.Users().AddUser(u), store.ErrUserIdDuplicated.Error())
}

func TestUserRepository_Find(t *testing.T) {
	s := New()
	u, ok := s.Users().FindByTgID(-1)
	assert.Nil(t, u)
	assert.Equal(t, ok, false)

	testUser := &model.User{IdTgbot: 1}
	s.Users().AddUser(testUser)
	u, ok = s.Users().FindByTgID(1)
	assert.Equal(t, u, testUser)
	assert.Equal(t, ok, true)
}
