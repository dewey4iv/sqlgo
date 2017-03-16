package main

import "github.com/dewey4iv/sqlgo"

func main() {
	db, err := sqlgo.New(
		sqlgo.WithDriver("postgres"),
		sqlgo.WithHostPort("127.0.0.1", "5432"),
	)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
}
