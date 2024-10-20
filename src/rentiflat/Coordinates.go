package rentiflat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type OSMResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetCoordinate(addressQuery string) (float64, float64, error) {
	addressQuery = strings.ReplaceAll(addressQuery, " ", "+")

	// Construct the request URL
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", addressQuery)

	// Send HTTP GET request to OSM API
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return 0, 0, fmt.Errorf(err.Error())
	}

	defer response.Body.Close()

	// Parse the JSON response
	var osmResponses []OSMResponse
	err = json.NewDecoder(response.Body).Decode(&osmResponses)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return 0, 0, fmt.Errorf(err.Error())
	}

	// Extract the coordinates from the first result
	if len(osmResponses) > 0 {
		lat, _ := strconv.ParseFloat(osmResponses[0].Lat, 64)
		lon, _ := strconv.ParseFloat(osmResponses[0].Lon, 64)
		fmt.Println("Latitude: ", lat, " Longitude: ", lon)
		return lat, lon, nil
	}
	return 0, 0, fmt.Errorf("coordinate not found")
}
