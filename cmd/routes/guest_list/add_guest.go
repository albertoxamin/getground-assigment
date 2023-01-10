package guest_list

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/getground/tech-tasks/backend/cmd/common/models"
)

type AddGuestRequestBody struct {
	Table       int `json:"table"`
	AccompanyingGuests       int `json:"accompanying_guests"`
}

func (h handler) AddGuest(c *gin.Context) {
	// Extract the request body as a Guest struct
    var req := AddGuestRequestBody{}
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Fetch the table from the database
    var table Table
    if err := h.DB.First(&table, req.TableID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid table"})
        return
    }

    // check if there is sufficient space at the table
    var count int
    if err := h.DB.Model(&Guest{}).Where("table_id = ?", req.TableID).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "problem counting guests"})
        return
    }
    if count+req.AccompanyingGuests > table.Capacity {
        c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient space"})
        return
    }

    // Add the guest to the guestlist
    guest := Guest{
        Table: req.TableID,
        AccompanyingGuests: req.AccompanyingGuests,
    }
    if err := h.DB.Create(&guest).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add guest"})
        return
    }

    // Return the name of the newly created guest
    c.JSON(http.StatusOK, gin.H{"name": guest.Name})
}