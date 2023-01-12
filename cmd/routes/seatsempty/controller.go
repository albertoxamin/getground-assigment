package seatsempty

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

// RegisterRoutes registers the routes for the seats_empty controller.
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/seats_empty")
	routes.GET("/", h.GetSeats)
}
