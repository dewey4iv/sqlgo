package sqlgo

import (
	"database/sql"
	"fmt"
	// import postgres
	// _ "github.com/lib/pq"
	// import mysql
	// _ "github.com/go-sql-driver/mysql"
)

// New takes the options and returns a DB instance
func New(opts ...Option) (DB, error) {
	var db db

	for _, opt := range opts {
		if err := opt.Apply(&db); err != nil {
			return nil, err
		}
	}

	db.opts = opts

	if err := verify(&db); err != nil {
		return nil, err
	}

	if err := db.Connect(); err != nil {
		return nil, err
	}

	return &db, nil
}

// db - exported for convenience
type db struct {
	opts              []Option
	conf              Config
	currentReconnects int64
	*sql.DB
}

// Database returns the underlying sql.DB
func (db *db) Database() *sql.DB {
	return db.DB
}

// Connect creates a connection to the DB
func (db *db) Connect() error {
	c := db.conf.Creds
	conn, err := sql.Open(db.conf.Driver, fmt.Sprintf(db.conf.ConnFmt, c.Host, c.Port, c.User, c.Password, c.DB))
	if err != nil {
		return err
	}

	db.DB = conn

	return nil
}

// Reconnect adds the ability to reconnect if there's an issue with the connection
func (db *db) Reconnect() error {
	return nil
}

// Config holds connection details
type Config struct {
	Driver        string `json:"driver"`
	ConnFmt       string `json:"conn_fmt"`
	MaxReconnects int64  `json:"max_reconnects"`
	Creds         Creds  `json:"creds"`
}

// Creds holds the fields for generic credentials
type Creds struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}
