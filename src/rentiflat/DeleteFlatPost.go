package rentiflat

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sachinmangla/rentiflat/database"
)

// @Summary Delete a flat post
// @Description Deletes a flat post by its ID
// @Tags flats
// @Accept json
// @Produce json
// @Param flat_id path int true "Flat ID"
// @Security ApiKeyAuth
// @Success 200 {string} string "Successfully deleted"
// @Failure 400 {string} string "Invalid flat ID"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Not Authorized to delete the given flat detail"
// @Failure 404 {string} string "Entry not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /flats/{flat_id} [delete]
func DeleteFlatPost(w http.ResponseWriter, r *http.Request) {
	var flatDetail database.FlatDetails

	// Parse flat_id from URL parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["flat_id"])
	if err != nil {
		http.Error(w, "Invalid flat ID", http.StatusBadRequest)
		return
	}

	// Fetch the flat details from the database
	log.Println("Fetching flat details for ID:", id)
	result := database.GetDb().Where("ID = ?", id).First(&flatDetail)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Entry not found", http.StatusNotFound)
			return
		}
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the user is authorized to delete the flat details
	userID := r.Context().Value("userID").(int)
	if userID != int(flatDetail.OwnerID) {
		http.Error(w, "Not Authorized to delete the given flat detail", http.StatusForbidden)
		return
	}

	// Delete the flat details from the database
	log.Println("Deleting flat details for ID:", id)
	result = database.GetDb().Unscoped().Delete(&flatDetail) // Use Unscoped() to permanently delete
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("Successfully deleted flat details for ID:", id)
}
