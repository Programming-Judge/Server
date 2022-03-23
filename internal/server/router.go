package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS_Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Authorization, Accept-Encoding, X-CSRF-Token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		//c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("OptionsPassthrough", "true")

		/*if c.Request.Method == "OPTIONS" {
			log.Print("got here")
			c.Status(200)
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

	problem := router.Group("/problem")
	{
		problem.POST("/create", Create)
		problem.GET("/view/:id", View)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
