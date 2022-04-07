package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Programming-Judge/Server/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Submit(ctx *gin.Context) {
	file, err := ctx.FormFile("code")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(file.Filename)
	uniqueName := uuid.New().String()
	newFileName := uniqueName + extension
	codeFile := strings.Replace(newFileName, "-", "", -1)
	fmt.Println(codeFile)
	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", codeFile)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	filename := uniqueName
	language := extension
	question := ctx.PostForm("QuestionID")
	user := ctx.PostForm("UserID")
	qsID, _ := strconv.Atoi(question)
	userID, _ := strconv.Atoi(user)

	qs, _ := store.FetchQuestion(qsID)
	tl := strconv.Itoa(qs.TimeLimit)
	ml := strconv.Itoa(qs.MemoryLimit)

	go SendEvaluator(question, filename, language, tl, ml, qsID, userID)
	fmt.Println(filename)
	fmt.Println(language)
	// File saved successfully. Return proper result
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

func SendEvaluator(question, filename, language, tl, ml string, qsID, userID int) {
	params := url.Values{}
	params.Add("id", question)
	params.Add("filename", filename)
	params.Add("lang", language)
	params.Add("timelimit", tl)
	params.Add("memorylimit", ml)
	body := strings.NewReader(params.Encode())

	req, _ := http.NewRequest("GET", "localhost:7070/submit/eval", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error in response from evaluator")
		return
	}
	defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//bodyString := string(b)
	status := 1

	if err := store.AddSubmission(filename, language, status, qsID, userID); err != nil {
		log.Print("Error in adding submission to database")
		return
	}
}
