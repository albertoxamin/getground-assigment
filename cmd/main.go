package main

import (
	"github.com/gin-gonic/gin"
	"github.com/getground/tech-tasks/backend/cmd/routes/tables"
	"github.com/getground/tech-tasks/backend/cmd/routes/guests"
	"github.com/getground/tech-tasks/backend/cmd/routes/guest_list"
	"github.com/getground/tech-tasks/backend/cmd/common/db"
)

func main() {
	port := "3000"
	dbUrl := "user:password@/getground"

	r := gin.Default()
	// h := db.Init(dbUrl)

	// tables.RegisterRoutes(r, h)
	print("fart sounds")

	r.Run(port)
}