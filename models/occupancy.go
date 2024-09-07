package models

import (
	"gorm.io/gorm"
	"time"
)

type Occupancy struct {
	ID                  uint      `gorm:"primaryKey"`
	RoomID              int       `gorm:"index"`
	OccupancyDate       time.Time
	OccupancyPercentage float64
}

// GetOccupancyPercentage retrieves the average occupancy percentage for a given room
// over the last 5 months. It uses a raw SQL query to compute the average from the 
// "occupancies" table, filtered by room ID and date range.
func GetOccupancyPercentage(db *gorm.DB, roomID int) (float64, error) {
	var avgOccupancy *float64

	query := `
		SELECT AVG(occupancy_percentage)
		FROM occupancy
		WHERE room_id = $1 AND occupancy_date >= CURRENT_DATE - INTERVAL '5 months'
	`

	err := db.Raw(query, roomID).Scan(&avgOccupancy).Error
	if err != nil {
		return 0, err
	}

	if avgOccupancy == nil {
		return 0, nil
	}

	return *avgOccupancy, nil
}