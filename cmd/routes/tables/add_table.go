package tables

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"github.com/gin-gonic/gin"
)

type addTableRequestBody struct {
	Capacity int `json:"capacity"`
}

func (h *handler) AddTable(c *gin.Context) {
	var body addTableRequestBody

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	println(body.Capacity)
	if body.Capacity < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "capacity must be greater than 0"})
		return
	}

	var table models.Table
	table.Capacity = body.Capacity

	if result := h.DB.Create(&table); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &table)
}
