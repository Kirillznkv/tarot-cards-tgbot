package mapstore

import (
	"github.com/stretchr/testify/assert"
	"tarot-cards-tgbot/internal/model"
	"tarot-cards-tgbot/internal/store"
	"testing"
)

func TestUserRepository_AddUser(t *testing.T) {
	s := New()
	u := &model.User{ID: 1}
	assert.NoError(t, s.Users().AddUser(u))
	assert.EqualError(t, s.Users().AddUser(u), store.ErrUserIdDuplicated.Error())
}

func TestUserRepository_Find(t *testing.T) {
	s := New()
	u, ok := s.Users().Find(-1)
	assert.Nil(t, u)
	assert.Equal(t, ok, false)

	testUser := &model.User{ID: 1}
	s.Users().AddUser(testUser)
	u, ok = s.Users().Find(1)
	assert.Equal(t, u, testUser)
	assert.Equal(t, ok, true)
}
