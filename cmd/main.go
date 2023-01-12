package main

import (
	"github.com/getground/tech-tasks/backend/cmd/common/db"
	"github.com/getground/tech-tasks/backend/cmd/routes/guestlist"
	"github.com/getground/tech-tasks/backend/cmd/routes/guests"
	"github.com/getground/tech-tasks/backend/cmd/routes/seatsempty"
	"github.com/getground/tech-tasks/backend/cmd/routes/tables"
	"github.com/gin-gonic/gin"
)

const port = ":3000"
const dbURL = "user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	r := gin.Default()
	h := db.Init(dbURL)

	// register routes
	tables.RegisterRoutes(r, h)
	guests.RegisterRoutes(r, h)
	guestlist.RegisterRoutes(r, h)
	seatsempty.RegisterRoutes(r, h)

	r.Run(port)
}
