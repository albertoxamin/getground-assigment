package main

import (
	"github.com/getground/tech-tasks/backend/cmd/common/db"
	"github.com/getground/tech-tasks/backend/cmd/routes/guest_list"
	"github.com/getground/tech-tasks/backend/cmd/routes/guests"
	"github.com/getground/tech-tasks/backend/cmd/routes/seats_empty"
	"github.com/getground/tech-tasks/backend/cmd/routes/tables"
	"github.com/gin-gonic/gin"
)

func main() {
	port := ":3000"
	dbUrl := "user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

	r := gin.Default()
	h := db.Init(dbUrl)

	tables.RegisterRoutes(r, h)
	guests.RegisterRoutes(r, h)
	guest_list.RegisterRoutes(r, h)
	seats_empty.RegisterRoutes(r, h)

	r.Run(port)
}
