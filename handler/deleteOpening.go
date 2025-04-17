package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/gopi/schemas"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening id"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {

		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id %s", id))
		return
	}

	sendSuccess(c, "delete-opening", opening)
}
