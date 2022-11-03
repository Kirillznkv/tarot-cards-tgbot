package mapstore

import (
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/model"
	"github.com/Kirillznkv/tarot-cards-tgbot/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_AddUser(t *testing.T) {
	s := New()
	u := &model.User{IdChat: 1, Name: "kshanti"}
	assert.NoError(t, s.Users().AddUser(u))
	assert.EqualError(t, s.Users().AddUser(u), store.ErrUserIdDuplicated.Error())
}

func TestUserRepository_Find(t *testing.T) {
	s := New()
	u, ok := s.Users().FindByTgID(-1)
	assert.Nil(t, u)
	assert.Equal(t, ok, false)

	testUser := &model.User{IdChat: 1, Name: "kshanti"}
	s.Users().AddUser(testUser)
	u, ok = s.Users().FindByTgID(1)
	assert.Equal(t, u, testUser)
	assert.Equal(t, ok, true)
}
