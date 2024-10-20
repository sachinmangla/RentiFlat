package rentiflat

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sachinmangla/rentiflat/database"
)

func checkUserExist(userId int) (database.OwnerDetails, error) {
	var owner database.OwnerDetails
	db := database.GetDb()

	result := db.Where("ID = ?", userId).First(&owner)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return database.OwnerDetails{}, errors.New("user-id not found")
		}
		log.Printf("Error fetching user with ID %d: %v", userId, result.Error)
		return database.OwnerDetails{}, result.Error
	}

	return owner, nil
}

// @Summary Create a new flat post
// @Description Create a new flat listing
// @Tags flats
// @Accept json
// @Produce json
// @Param flat body database.FlatDetails true "Flat details"
// @Param Authorization header string true "Bearer {token}"
// @Success 201 {object} database.FlatDetails
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /add-user [post]
func RentiFlatCreatePost(w http.ResponseWriter, r *http.Request) {
	var flat database.FlatDetails

	err := json.NewDecoder(r.Body).Decode(&flat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value("userID").(int)
	owner, err := checkUserExist(userID)

	if err != nil {
		http.Error(w, "User-ID not found or incorrect", http.StatusBadRequest)
		return
	}
	flat.OwnerID = uint(userID)
	flat.Owner = owner

	lat, lon, err := GetCoordinate(flat.Address)
	fmt.Print(lat, lon)

	if err != nil {
		http.Error(w, "Address not correct", http.StatusBadRequest)
		return
	}

	flat.Location = database.Coordinates{Latitude: lat, Longitude: lon}

	db := database.GetDb()
	result := db.Create(&flat)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(flat)
}
