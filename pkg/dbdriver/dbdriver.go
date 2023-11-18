package dbdriver

import (
	"database/sql"
	"time"

	// Install PGX driver for sql library
	// go get github.com/jackc/pgx/v4

	// In directory pkg/dbdriver
	// NEW Driver file that is setup to use Postgres
	// as a database, but can be changed to work with
	// any database

	// Low level access to Postgres
	// _ "github.com/jackc/pgx/pgconn"

	_ "github.com/jackc/pgx/v4"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Holds database connection
// This is where you can define you want to use
// a different type of database
type DB struct {
	SQL *sql.DB
}

// Initialize a reference to DB type
var dbConn = &DB{}

// Max number of connections allowed to DB
const maxOpenDbConns = 20

// How many idle connections will we allow
const maxIdleDbConns = 10

// Max lifetime for DB connection (5 minutes)
const maxDbLifetime = 5 * time.Minute

// Try to create database
// Pass a database connection string
// Returns pointer to DB and potential error
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Create a connection pool
// Pass database connection string
// Returns pointer to DB type and error
func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)

	// If I can't connect to the DB die
	if err != nil {
		panic(err)
	}

	// If it does work set parameters
	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetMaxIdleConns(maxIdleDbConns)
	db.SetConnMaxLifetime(maxDbLifetime)

	// Set DB to dbConn
	dbConn.SQL = db

	return dbConn, nil
}
