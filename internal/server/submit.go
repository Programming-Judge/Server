package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Programming-Judge/Server/internal/store"
	"github.com/gin-gonic/gin"
)

func Submit(ctx *gin.Context) {
	filename := ctx.PostForm("FileName")
	language := ctx.PostForm("Language")
	question := ctx.PostForm("QuestionID")
	user := ctx.PostForm("UserID")
	qsID, _ := strconv.Atoi(question)
	userID, _ := strconv.Atoi(user)

	qs, _ := store.FetchQuestion(qsID)
	tl := strconv.Itoa(qs.TimeLimit)
	ml := strconv.Itoa(qs.MemoryLimit)

	go SendEvaluator(question, filename, language, tl, ml, qsID, userID)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Successfully Submitted",
	})
}

func SendEvaluator(question, filename, language, tl, ml string, qsID, userID int) {
	params := url.Values{}
	params.Add("id", question)
	params.Add("filename", filename)
	params.Add("language", language)
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
	b, err := ioutil.ReadAll(resp.Body)
	bodyString := string(b)
	status := 1

	if err := store.AddSubmission(filename, language, status, qsID, userID); err != nil {
		log.Print("Error in adding submission to database")
		return
	}
}
