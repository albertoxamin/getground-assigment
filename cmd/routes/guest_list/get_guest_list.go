package guest_list

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetGuestList(c *gin.Context) {
	var guests []models.Guest

	if result := h.DB.Find(&guests); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &guests)
}
