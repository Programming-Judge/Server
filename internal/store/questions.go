package store

import (
	"log"
)

type Question struct {
	ID          int
	Title       string `binding:"required,min=3,max=20"`
	Description string `binding:"required,min=10,max=200"`
	TimeLimit   int
	MemoryLimit int
}

func AddQuestion(qs *Question) error {
	_, err := db.Model(qs).Returning("*").Insert()
	if err != nil {
		log.Printf("Error inserting new post")
	}
	return err
}

func FetchQuestion(id int) (*Question, error) {
	qs := new(Question)
	qs.ID = id
	err := db.Model(qs).WherePK().Select()
	if err != nil {
		log.Printf("Error fetching post")
		return nil, err
	}
	return qs, nil
}