package data

// Package data handles all database interactions and data models

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Models wraps all database models for the application
type Models struct {
	Movies MovieModel // Handles all movie-related database operations
}

// NewModels creates a new Models instance with the provided database connection
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
