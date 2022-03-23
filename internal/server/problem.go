package server

import (
	"net/http"
	"strconv"

	"github.com/Programming-Judge/Server/internal/store"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	qs := new(store.Question)
	if err := ctx.Bind(qs); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := store.AddQuestion(qs); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
		"msg":  "Posts fetched successfully.",
		"data": qs,
	})
}
