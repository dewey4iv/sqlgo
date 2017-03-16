package sqlgo

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// WithJSONConfigBytes takes A JSON Object as a []byte sets its config options
func WithJSONConfigBytes(b []byte) Option {
	return &withJSONConfig{b}
}

type withJSONConfig struct {
	b []byte
}

func (opt *withJSONConfig) Apply(db *db) error {
	if err := json.Unmarshal(opt.b, &db.conf); err != nil {
		return OptErr("WithJSONConfig", err)
	}

	return nil
}

// WithCreds sets the user, password, and database
func WithCreds(user string, password string, database string) Option {
	return &withCreds{user, password, database}
}

type withCreds struct {
	user     string
	password string
	database string
}

func (opt *withCreds) Apply(db *db) error {
	db.conf.Creds.User = opt.user
	db.conf.Creds.Password = opt.password
	db.conf.Creds.DB = opt.database

	return nil
}

// WithHostPort takes the host and port for connecting to the database
func WithHostPort(host string, port string) Option {
	return &withHostPort{host, port}
}

type withHostPort struct {
	host string
	port string
}

func (opt *withHostPort) Apply(db *db) error {
	db.conf.Creds.Host = opt.host
	db.conf.Creds.Port = opt.port

	return nil
}

// WithSRVRecord takes parts for an SRV record and connects to an available server
func WithSRVRecord(service string, proto string, name string) Option {
	return &withSRVRecord{service, proto, name}
}

type withSRVRecord struct {
	service string
	proto   string
	name    string
}

func (opt *withSRVRecord) Apply(db *db) error {
	return nil
}

// WithCreds

// WithDBConn takes a sql.DB and sets it directly
func WithDBConn(db *sql.DB) Option {
	return &withDBConn{db}
}

type withDBConn struct {
	db *sql.DB
}

func (opt *withDBConn) Apply(db *db) error {
	db.DB = opt.db

	return nil
}

// WithReconnectStrategy

// WithDriver takes the driver name as a string and sets it
func WithDriver(driver string) Option {
	return &withDriver{driver}
}

type withDriver struct {
	driver string
}

func (opt *withDriver) Apply(db *db) error {
	switch opt.driver {
	case "postgres":
		db.conf.Driver = opt.driver
		db.conf.ConnFmt = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	case "mysql":
		db.conf.Driver = opt.driver
		db.conf.ConnFmt = "%s:%s@tcp(%s:%s)/%s"
	default:
		return fmt.Errorf(`driver "%s" not supported`, opt.driver)
	}

	return nil
}

func WithConfig(config Config) Option {
	return &withConfig{config}
}

type withConfig struct {
	config Config
}

func (opt *withConfig) Apply(db *db) error {
	db.conf = opt.config

	return nil
}
