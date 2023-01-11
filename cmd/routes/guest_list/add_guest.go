package guest_list

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"github.com/gin-gonic/gin"
)

type AddGuestRequestBody struct {
	Table              int `json:"table"`
	AccompanyingGuests int `json:"accompanying_guests"`
}

func (h handler) AddGuest(c *gin.Context) {
	// Extract the request body as a Guest struct
	var req = AddGuestRequestBody{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the table from the database
	var table models.Table
	if err := h.DB.First(&table, req.Table).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid table"})
		return
	}

	// check if there is sufficient space at the table
	var count int64
	if err := h.DB.Model(&models.Guest{}).Where("`table` = ?", req.Table).Select("count(*)+ifnull(sum(accompanying_guests), 0)").Row().Scan(&count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem counting guests"})
		return
	}

	if count+int64(req.AccompanyingGuests)+1 > int64(table.Capacity) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient space"})
		return
	}

	// Add the guest to the guestlist
	guest := models.Guest{
		Name:               c.Param("name"),
		Table:              req.Table,
		AccompanyingGuests: req.AccompanyingGuests,
	}
	if err := h.DB.Create(&guest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add guest"})
		return
	}

	// Return the name of the newly created guest
	c.JSON(http.StatusOK, gin.H{"name": guest.Name})
}
