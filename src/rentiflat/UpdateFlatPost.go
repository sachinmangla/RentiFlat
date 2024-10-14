package rentiflat

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sachinmangla/rentiflat/database"
)

func UpdateFlatDetail(w http.ResponseWriter, r *http.Request) {
	var flatDetail database.FlatDetails
	var updatedFlatDetail database.UpdateFlatDetail

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["flat_id"])

	result := database.GetDb().Where("ID = ?", id).First(&flatDetail)

	if result.Error != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	userID := r.Context().Value("userID").(int)

	if userID != int(flatDetail.OwnerID) {
		http.Error(w, "Not Authorized to update the given flat detail", http.StatusForbidden)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updatedFlatDetail)

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

	result = database.GetDb().Save(&flatDetail)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flatDetail)
}
