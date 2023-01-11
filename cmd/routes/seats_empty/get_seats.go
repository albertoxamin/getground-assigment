package seats_empty

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetSeats(c *gin.Context) {
	var guest_count int64
	if err := h.DB.Raw("select ifnull(sum(accompanying_guests)+count(id), 0) from guests where deleted_at is NULL;").Row().Scan(&guest_count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem counting guests"})
		return
	}
	var seats_count int64
	if err := h.DB.Model(&models.Table{}).Select("sum(capacity)").Row().Scan(&seats_count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem counting seats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seats_empty": seats_count - guest_count})
}
