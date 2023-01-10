package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Capacity int `json:"capacity"`
}