package model

import "database/sql"

type Model struct {
	DB *sql.DB
}

// TODO(tho) pass the db
func NewModel(db *sql.DB) *Model {
	return &Model{
		DB: db,
	}
}
