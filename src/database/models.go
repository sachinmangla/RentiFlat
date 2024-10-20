package database

import (
	"github.com/jinzhu/gorm"
)

// OwnerDetails represents the details of an owner
type OwnerDetails struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" gorm:"type:varchar(100);not null"`         // Added type for better database control
	Email      string `json:"email" gorm:"type:varchar(100);not null;unique"` // Added type and unique constraints
	Phone      string `json:"phone" gorm:"type:varchar(20);not null"`         // Added type for better database control
	Password   string `json:"password" gorm:"type:varchar(100);not null"`
}

// Coordinates represents latitude and longitude values
type Coordinates struct {
	Latitude  float64 `json:"latitude"`  // Latitude value
	Longitude float64 `json:"longitude"` // Longitude value
}

// Note: Ensure the necessary PostgreSQL extensions are enabled
// CREATE EXTENSION cube;
// CREATE EXTENSION earthdistance;

// FlatDetails represents the details of a flat
type FlatDetails struct {
	gorm.Model      `swaggerignore:"true"`
	OwnerID         uint         `json:"owner_id" gorm:"not null" swaggerignore:"true"`                                                       // Foreign key to OwnerDetails
	Owner           OwnerDetails `json:"owner" swaggerignore:"true" gorm:"foreignkey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Owner relationship
	Location        Coordinates  `json:"location" gorm:"embedded;embeddedPrefix:location_"`                                                   // Embedded coordinates
	Address         string       `json:"address" gorm:"type:varchar(255);not null"`                                                           // Address of the flat
	Rent            float64      `json:"rent" gorm:"type:decimal(10,2);not null"`                                                             // Rent amount
	SecurityDeposit float64      `json:"security_deposit" gorm:"type:decimal(10,2);not null"`                                                 // Security deposit amount
	LookingFor      string       `json:"looking_for" gorm:"type:varchar(100);not null"`                                                       // Target tenant description
}

type LoginDetail struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type UpdateFlatDetail struct {
	Address         string  `json:"address,omitempty"`          // Address of the flat
	Rent            float64 `json:"rent,omitempty"`             // Rent amount
	SecurityDeposit float64 `json:"security_deposit,omitempty"` // Security deposit amount
	LookingFor      string  `json:"looking_for,omitempty"`
}

type Response struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
}

type JwtToken struct {
	Token string `json:"token"`
}
