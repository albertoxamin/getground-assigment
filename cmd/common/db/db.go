package db

import (
	"log"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init initializes the database connection and returns a pointer to the database managed by gorm.
func Init(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Here we register the models to use with gorm.
	// This is done by calling AutoMigrate on the database connection.
	// Which will create the tables if they don't exist.
	// If they do exist, it will update the tables to match the models.
	db.AutoMigrate(&models.Table{})
	db.AutoMigrate(&models.Guest{})

	return db
}
