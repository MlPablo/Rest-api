package teststore

import (
	"github.com/MlPablo/rest-API/internal/app/model"
	"github.com/MlPablo/rest-API/internal/app/store"
)

// Store ...
type Store struct {
	userRepositoty *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepositoty != nil {
		return s.userRepositoty
	}

	s.userRepositoty = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepositoty
}
