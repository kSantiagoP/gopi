// Cada nova pasta, assume-se um submodule
// Normalmente cada pasta vai ter um arquivo .go de nome igual dentro
package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	//Definir rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":4000") // listen and serve on localhost:4000
}
