package server

import (
	"log"
	"net/http"

	"github.com/Programming-Judge/Server/internal/store"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := store.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": "123456789",
	})
}

func Login(ctx *gin.Context) {
	user := new(store.User)
	ctx.Request.ParseForm()

	username := ctx.PostForm("username")
	pass := ctx.PostForm("password")

	user.Username = username
	user.Password = pass

	log.Printf(user.Username)
	log.Printf(user.Password)
	/*	if user, err := ctx.FormFile(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} */

	user, err := store.Authenticate(user.Username, user.Password)
	if err != nil {
		log.Print("here i am")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": "123456789",
	})
}
