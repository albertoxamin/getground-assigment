package tables

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

// RegisterRoutes registers the routes for the tables controller.
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/tables")
	routes.POST("/", h.AddTable)
}
