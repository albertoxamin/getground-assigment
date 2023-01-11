package guests

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/guests")
	routes.GET("/", h.GetCheckedGuestList)
	routes.PUT("/:name", h.CheckInGuest)
	routes.DELETE("/:name", h.CheckOutGuest)
}
