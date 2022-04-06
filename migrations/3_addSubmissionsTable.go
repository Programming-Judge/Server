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
      	status INT,
      	language TEXT NOT NULL,
	  	user_id INT REFERENCES users ON DELETE CASCADE,
	  	question_id INT REFERENCES questions ON DELETE CASCADE
    )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table submissions...")
		_, err := db.Exec(`DROP TABLE submissions`)
		return err
	})
}
