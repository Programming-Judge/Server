package store

import "log"

type Submission struct {
	ID         int
	FileName   string
	UserName   string
	Status     int
	Language   string
	QuestionID int
}

func AddSubmission(filename, language, username string, status, qsID int) error {
	sub := &Submission{
		FileName:   filename,
		UserName:   username,
		Status:     status,
		Language:   language,
		QuestionID: qsID,
	}

	_, err := db.Model(sub).Returning("*").Insert()
	if err != nil {
		log.Printf("Error inserting new submission" , err)
	}
	return err
}
