package seatsempty

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetSeats(c *gin.Context) {
	var guestCount int64
	if err := h.DB.Raw("select ifnull(sum(accompanying_guests)+count(id), 0) from guests where deleted_at is NULL;").Row().Scan(&guestCount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem counting guests"})
		return
	}
	var seatsCount int64
	if err := h.DB.Model(&models.Table{}).Select("ifnull(sum(capacity), 0)").Row().Scan(&seatsCount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem counting seats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seats_empty": seatsCount - guestCount})
}
