package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Programming-Judge/Server/internal/store"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	title := ctx.PostForm("title")
	desc := ctx.PostForm("description")
	ml := ctx.PostForm("memorylimit")
	tl := ctx.PostForm("timelimit")
	mlval, _ := strconv.Atoi(ml)
	tlval, _ := strconv.Atoi(tl)

	qs := new(store.Question)
	qs.Title = title
	qs.Description = desc
	qs.MemoryLimit = mlval
	qs.TimeLimit = tlval
	/*if err := ctx.Bind(qs); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/
	if err := store.AddQuestion(qs); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//fmt.Println(qs.ID)
	inputzip, err := ctx.FormFile("input")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outputzip, err := ctx.FormFile("output")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	namingID := strconv.Itoa(qs.ID)
	nameinzip := "inzip" + namingID
	if err := ctx.SaveUploadedFile(inputzip, fmt.Sprintf("../Storage/tests/%s", nameinzip)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	nameoutzip := "outzip" + namingID
	if err := ctx.SaveUploadedFile(outputzip, fmt.Sprintf("../Storage/tests/%s", nameoutzip)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	finalIN := namingID + "-input"
	finalOUT := namingID + "-output"
	err1 := Unzip("../Storage/tests/"+nameinzip, "../Storage/tests/"+finalIN)
	if err1 != nil {
		log.Println("could not unzip input cases")
	}
	err2 := Unzip("../Storage/tests/"+nameoutzip, "../Storage/tests/"+finalOUT)
	if err2 != nil {
		log.Println("could not unzip output cases")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Added to Problemset.",
		"data": qs,
	})
}

func View(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	qs, err := store.FetchQuestion(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Problem fetched successfully.",
		"data": qs,
	})
}

func Update(ctx *gin.Context) {
	paramID := ctx.Param("id")
	title := ctx.PostForm("Title")
	description := ctx.PostForm("Description")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}

	if err := store.UpdateQuestion(id, title, description); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Problem Updated",
	})
}

func Delete(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}

	if err := store.DeleteQuestion(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Deleted from Problemset.",
	})
}

func ViewAll(ctx *gin.Context) {

	questions, err := store.FetchQuestions()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Problems fetched successfully.",
		"data": questions,
	})
}
