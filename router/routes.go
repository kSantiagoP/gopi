package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", ping)
		v1.GET("/opening", handleOpening)
		v1.POST("/opening", handleOpening)
		v1.DELETE("/opening", handleOpening)
		v1.PUT("/opening", handleOpening)
		v1.GET("/openings", handleOpening)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "The tcheca",
	})
}
