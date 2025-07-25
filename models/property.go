package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Location    string  `json:"location"`
}
