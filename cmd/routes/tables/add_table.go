package tables

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/getground/tech-tasks/backend/cmd/common/models"
)

type AddTableRequestBody struct {
	Capacity       string `json:"capacity"`
}

func (h handler) AddTable(c *gin.Context) {
	body := AddTableRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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