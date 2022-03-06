package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS_Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		/*
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}*/

		c.Next()
	}
}

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()
	router.Use(CORS_Middleware())

	auth := router.Group("/auth")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)
	}

	router.NoRoute(func(ctx *gin.Context) {
		log.Print("here")
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
