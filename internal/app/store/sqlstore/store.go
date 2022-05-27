package sqlstore

import (
	"database/sql"
	"github.com/MlPablo/rest-API/internal/app/store"
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepositoty *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{db: db}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepositoty != nil {
		return s.userRepositoty
	}

	s.userRepositoty = &UserRepository{
		store: s,
	}
	return s.userRepositoty
}
