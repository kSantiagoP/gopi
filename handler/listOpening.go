package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/gopi/schemas"
)

func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "internal error fetching openings")
		return
	}

	sendSuccess(c, "list-openings", openings)
}
