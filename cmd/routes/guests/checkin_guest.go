package guests

import (
	"net/http"
	"time"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"github.com/gin-gonic/gin"
)

type CheckInRequestBody struct {
	AccompanyingGuests int `json:"accompanying_guests"`
}

func (h *handler) CheckInGuest(c *gin.Context) {
	// Extract the request body as a Guest struct
	var req = CheckInRequestBody{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.AccompanyingGuests < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accompanying_guests must be at least 0"})
	}

	// Fetch the table from the database
	var table models.Table
	if err := h.DB.Model(&models.Guest{}).Where("`name` = ?", c.Param("name")).Joins("JOIN tables as t on t.id=`guests`.table").Select("t.*").First(&table).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid table"})
		return
	}

	var guest models.Guest
	h.DB.Model(&models.Guest{}).Where("`name` = ?", c.Param("name")).First(&guest)

	// check if there is sufficient space at the table
	var count int64
	if err := h.DB.Model(&models.Guest{}).Where("`table` = ?", table.ID).Where("`name` <> ?", c.Param("name")).Select("count(*)+ifnull(sum(accompanying_guests), 0)").Row().Scan(&count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem counting guests"})
		return
	}

	if count+int64(req.AccompanyingGuests)+1 > int64(table.Capacity) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient space"})
		return
	}
	guest.AccompanyingGuests = req.AccompanyingGuests
	guest.TimeArrived = time.Now().Format(time.RFC3339)

	if err := h.DB.Save(&guest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed check in guest"})
		return
	}

	// Return the name of the updated guest
	c.JSON(http.StatusOK, gin.H{"name": guest.Name})
}
