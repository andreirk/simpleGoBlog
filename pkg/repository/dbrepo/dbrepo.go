package dbrepo

// In pkg/repository/dbrepo

import (
	"database/sql"
	"web3/pkg/config"
	"web3/pkg/repository"
)

// Holds reference to our AppConfig
// that holds our session and CSRFToken
// It also holds our DB connection pool
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// Receives DB connection and AppConfig
// and returns a Postgres connection pool
// We can create other functions that work
// with other DBs
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	// Return a reference
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

// Now update main to initialize the
// DB repo
