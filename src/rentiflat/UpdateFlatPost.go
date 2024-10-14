package rentiflat

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sachinmangla/rentiflat/database"
)

// @Summary Update flat details
// @Description Update details of a specific flat
// @Tags flats
// @Accept json
// @Produce json
// @Param flat_id path int true "Flat ID"
// @Param updatedFlatDetail body database.UpdateFlatDetail true "Updated flat details"
// @Success 200 {object} database.FlatDetails
// @Failure 400 {string} string "Bad Request"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Security ApiKeyAuth
// @Router /flats/{flat_id} [put]
func UpdateFlatDetail(w http.ResponseWriter, r *http.Request) {
	var flatDetail database.FlatDetails
	var updatedFlatDetail database.UpdateFlatDetail

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["flat_id"])
	if err != nil {
		http.Error(w, "Invalid flat ID", http.StatusBadRequest)
		return
	}

	log.Println("Fetching flat details for ID:", id)
	result := database.GetDb().Where("ID = ?", id).First(&flatDetail)
	if result.Error != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	userID := r.Context().Value("userID").(int)
	if userID != int(flatDetail.OwnerID) {
		http.Error(w, "Not authorized to update the given flat detail", http.StatusForbidden)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updatedFlatDetail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lat, lon, err := GetCoordinate(updatedFlatDetail.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	flatDetail.Location.Latitude = lat
	flatDetail.Location.Longitude = lon
	flatDetail.Address = updatedFlatDetail.Address
	flatDetail.Rent = updatedFlatDetail.Rent
	flatDetail.SecurityDeposit = updatedFlatDetail.SecurityDeposit
	flatDetail.LookingFor = updatedFlatDetail.LookingFor

	log.Println("Updating flat details:", flatDetail)
	result = database.GetDb().Save(&flatDetail)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flatDetail)
}
