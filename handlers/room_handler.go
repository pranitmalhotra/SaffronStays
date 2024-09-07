package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"saffronstays-api/models"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
)

type RoomResponse struct {
	RoomID              int     `json:"room_id"`
	OccupancyPercentage float64 `json:"occupancy_percentage"`
	AvgRate             float64 `json:"average_night_rate"`
	HighRate            float64 `json:"highest_night_rate"`
	LowRate             float64 `json:"lowest_night_rate"`
}

// RoomHandler returns an http.HandlerFunc that processes requests to the /api/room/{room_id} endpoint.
// It retrieves the room ID from the URL path, fetches occupancy and rate data for the specified room,
// and returns this data as a JSON response. In case of errors, it returns appropriate HTTP error responses.
func RoomHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomID, err := strconv.Atoi(vars["room_id"])
		if err != nil {
			http.Error(w, "Invalid room_id", http.StatusBadRequest)
			return
		}

		occupancy, err := models.GetOccupancyPercentage(db, roomID)
		if err != nil {
			http.Error(w, "Error fetching occupancy data", http.StatusInternalServerError)
			return
		}

		avgRate, highRate, lowRate, err := models.GetRates(db, roomID)
		if err != nil {
			http.Error(w, "Error fetching rate data", http.StatusInternalServerError)
			return
		}

		if occupancy == 0 && avgRate == 0 && highRate == 0 && lowRate == 0 {
			http.Error(w, "Data is available only for Room 1 & 2", http.StatusNotFound)
			return
		}

		response := RoomResponse{
			RoomID:              roomID,
			OccupancyPercentage: occupancy,
			AvgRate:             avgRate,
			HighRate:            highRate,
			LowRate:             lowRate,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}