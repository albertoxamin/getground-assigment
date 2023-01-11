package db

import (
	"log"

	"github.com/getground/tech-tasks/backend/cmd/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Table{})
	db.AutoMigrate(&models.Guest{})

	return db
}
