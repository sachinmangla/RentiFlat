package rentiflat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sachinmangla/rentiflat/database"
)

// @Summary Search for flats
// @Description Search for flats within a specified radius of a given address
// @Tags flats
// @Accept json
// @Produce json
// @Param q query string true "Address to search around"
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {array} database.FlatDetails
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /search [get]
func SearchFlat(w http.ResponseWriter, r *http.Request) {
	var flats []database.FlatDetails
	address := r.URL.Query().Get("q")
	radius := 10000

	fmt.Println("address is :- ", address)
	if address == "" {
		http.Error(w, "Query empty in the Request", http.StatusBadRequest)
	}
	latitude, longitude, err := GetCoordinate(address)
	if err != nil {
		http.Error(w, "Address not correct", http.StatusBadRequest)
	}

	query := `
        SELECT fd.*, od.*
        FROM flat_details fd
        JOIN owner_details od ON fd.owner_id = od.id
        WHERE earth_box(ll_to_earth(?, ?), ?) @> ll_to_earth(fd.latitude, fd.longitude);
    `

	db := database.GetDb()

	if err := db.Raw(query, latitude, longitude, radius).Scan(&flats).Error; err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flats)
}
