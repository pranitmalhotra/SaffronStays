package models

import (
	"gorm.io/gorm"
)

type Rate struct {
	ID        uint    `gorm:"primaryKey"`
	RoomID    int     `gorm:"column:room_id"`
	RateDate  string  `gorm:"column:rate_date"`
	NightRate float64 `gorm:"column:night_rate"`
}

// Explicitly specify the table name to be "rate"
func (Rate) TableName() string {
	return "rate"
}

type RateStats struct {
	AvgRate  float64 `gorm:"column:avg_rate"`
	HighRate float64 `gorm:"column:max_rate"`
	LowRate  float64 `gorm:"column:min_rate"`
}

// GetRates retrieves the average, highest, and lowest night rates for a given room
// within the next 30 days. It calculates these statistics using a query on the "rate" 
// table filtered by room ID and date range.
func GetRates(db *gorm.DB, roomID int) (float64, float64, float64, error) {
	var stats RateStats
	err := db.Model(&Rate{}).
		Where("room_id = ? AND rate_date BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '30 days'", roomID).
		Select("AVG(night_rate) as avg_rate, MAX(night_rate) as max_rate, MIN(night_rate) as min_rate").
		Scan(&stats).Error
	if err != nil {
		return 0, 0, 0, err
	}

	return stats.AvgRate, stats.HighRate, stats.LowRate, nil
}
