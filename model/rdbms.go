package model

import "database/sql"

// Query database/sql compatible query interface
type Query interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// Tx represents database transaction interface
type Tx interface {
	Query
	Commit() error
	Rollback() error
}

// DB database/sql interface
type DB interface {
	Query
	Begin() (*sql.Tx, error)
	Close() error
	Ping() error
}
