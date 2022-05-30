package teststore_test

import (
	"github.com/MlPablo/rest-API/internal/app/model"
	"github.com/MlPablo/rest-API/internal/app/store"
	"github.com/MlPablo/rest-API/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepositoty_Find(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	s.User().Create(u)

	u1, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u1)

}
