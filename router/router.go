// each new directory usually means a new submodule
// also, it usually (not always) have a .go code with the same name
package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//initializing router
	router := gin.Default()

	initializeRoutes(router)

	router.Run(":4000") // listen and serve on localhost:4000
}
