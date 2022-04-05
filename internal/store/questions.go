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
		log.Printf("Error inserting new question")
	}
	return err
}

func FetchQuestion(id int) (*Question, error) {
	qs := new(Question)
	qs.ID = id
	err := db.Model(qs).WherePK().Select()
	if err != nil {
		log.Printf("Error fetching question")
		return nil, err
	}
	return qs, nil
}

func FetchQuestions() ([]Question, error) {
	var questions []Question
	err := db.Model(&questions).Select()
	if err != nil {
		log.Printf("Error fetching questions")
		return nil, err
	}
	return questions, nil
}

func UpdateQuestion(id int, title, description string) error {
	qs := new(Question)
	qs.ID = id
	err := db.Model(qs).WherePK().Select()
	if err != nil {
		log.Printf("Error: Question does not exist")
		return err
	}
	qs.Description = description
	qs.Title = title
	res, err := db.Model(qs).WherePK().UpdateNotZero()
	if err != nil {
		log.Printf("Error: Could not Update question")
		return err
	}
	log.Println("Updated ", res.RowsAffected(), "row(s)")
	return nil
}

func DeleteQuestion(id int) error {
	qs := new(Question)
	qs.ID = id

	//check if question is present
	err := db.Model(qs).WherePK().Select()
	if err != nil {
		log.Printf("Error: Question does not exist")
		return err
	}

	//delete question
	res, err := db.Model(qs).WherePK().Delete()
	if err != nil {
		log.Printf("Error: Could not delete question")
		return err
	}

	//print number of rows affected
	log.Println("Deleted ", res.RowsAffected(), "row(s)")
	return nil
}
