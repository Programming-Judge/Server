package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table questions...")
		_, err := db.Exec(`CREATE TABLE questions(
      	id SERIAL PRIMARY KEY,
      	title TEXT NOT NULL,
      	description TEXT NOT NULL,
	  	time_limit INT,
	  	memory_limit INT
    )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table questions...")
		_, err := db.Exec(`DROP TABLE questions`)
		return err
	})
}
