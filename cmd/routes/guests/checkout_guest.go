package guests

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CheckOutGuest(c *gin.Context) {
	var guest models.Guest
	if err := h.DB.Model(&models.Guest{}).Where("`name` = ?", c.Param("name")).First(&guest).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Using the deleted_at column to mark the guest as checked out
	if err := h.DB.Delete(&guest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed check out guest"})
		return
	}

	// Return the name of the updated guest
	c.JSON(204, gin.H{})
}
