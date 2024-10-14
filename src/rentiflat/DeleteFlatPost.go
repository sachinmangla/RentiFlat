package rentiflat

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sachinmangla/rentiflat/database"
)

func DeleteFlatPost(w http.ResponseWriter, r *http.Request) {
	var flatDetail database.FlatDetails

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["flat_id"])

	result := database.GetDb().Where("ID = ?", id).First(&flatDetail)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Entry not found", http.StatusNotFound)
			return
		}
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.Context().Value("userID").(int)

	if userID != int(flatDetail.OwnerID) {
		http.Error(w, "Not Authorized to update the given flat detail", http.StatusForbidden)
		return
	}

	result = database.GetDb().Delete(&flatDetail)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
