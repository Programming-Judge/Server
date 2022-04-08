package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table submissions...")
		_, err := db.Exec(`CREATE TABLE submissions(
      	id SERIAL PRIMARY KEY,
		file_name TEXT NOT NULL,
		question_name TEXT NOT NULL,
		user_name TEXT NOT NULL,
      	status TEXT NOT,
      	language TEXT NOT NULL
    )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table submissions...")
		_, err := db.Exec(`DROP TABLE submissions`)
		return err
	})
}
