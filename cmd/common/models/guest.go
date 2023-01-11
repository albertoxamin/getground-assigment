package models

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	Name               string `json:"name"`
	Table              int    `json:"table"`
	AccompanyingGuests int    `json:"accompanying_guests"`
	TimeArrived        string `json:"time_arrived"`
}
