package database

import "github.com/jinzhu/gorm"

// Coordinates represents latitude and longitude values
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// FlatDetails represents details of a flat
type FlatDetails struct {
	gorm.Model
	Name            string      `gorm:"not null"`
	Email           string      `gorm:"not null"`
	Phone           string      `gorm:"not null"`
	Location        Coordinates `gorm:"not null"`
	Address         string
	Rent            float64
	SecurityDeposit float64
	LookingFor      string
}
