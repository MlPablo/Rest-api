package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepositoty *UserRepositoty
}

// New ...
func New(config *Config) *Store {
	return &Store{config: config}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(" +
		"id bigserial not null primary key," +
		"email varchar not null unique," +
		"encrypted_password varchar not null" +
		");")
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepositoty {
	if s.userRepositoty != nil {
		return s.userRepositoty
	}

	s.userRepositoty = &UserRepositoty{
		store: s,
	}
	return s.userRepositoty
}
