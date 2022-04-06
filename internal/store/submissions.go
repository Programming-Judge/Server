package store

import "log"

type Submission struct {
	ID         int
	FileName   string
	Status     int
	Language   string
	UserID     int
	QuestionID int
}

func AddSubmission(filename, language string, status, qsID, userID int) error {
	sub := Submission{
		FileName:   filename,
		Status:     status,
		Language:   language,
		UserID:     userID,
		QuestionID: qsID,
	}

	_, err := db.Model(sub).Returning("*").Insert()
	if err != nil {
		log.Printf("Error inserting new submission")
	}
	return err
}
