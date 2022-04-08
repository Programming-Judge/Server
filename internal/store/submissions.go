package store

import "log"

type Submission struct {
	ID           int
	FileName     string
	UserName     string
	QuestionName string
	Status       string
	Language     string
}

func AddSubmission(filename, language, username, status, qsName string) error {
	sub := &Submission{
		FileName:     filename,
		UserName:     username,
		QuestionName: qsName,
		Status:       status,
		Language:     language,
	}

	_, err := db.Model(sub).Returning("*").Insert()
	if err != nil {
		log.Print("Error inserting new submission", err)
	}
	return err
}

func FetchSubmissions() ([]Submission, error) {
	var submissions []Submission
	err := db.Model(&submissions).Select()
	if err != nil {
		log.Printf("Error fetching questions")
		return nil, err
	}
	return submissions, nil
}
