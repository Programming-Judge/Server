package server

import (
	"net/http"

	"github.com/Programming-Judge/Server/internal/store"
	"github.com/gin-gonic/gin"
)

func History(ctx *gin.Context) {
	submissions, err := store.FetchSubmissions()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Problems fetched successfully.",
		"data": submissions,
	})
}
