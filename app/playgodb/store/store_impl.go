package store

import "database/sql"

// SQLStore ...
type SQLStore struct {
	db *sql.DB
}

// NewStore ...
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{db: db}
}

// Ping ...
func (s *SQLStore) Ping() error {
	return s.db.Ping()
}

// Close ...
func (s *SQLStore) Close() error {
	return s.db.Close()
}
