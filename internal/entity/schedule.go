package entity

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	MovieTitle  string         `json:"movie_title"`
	CinemaName  string         `json:"cinema_name"`
	StudioName  string         `json:"studio_name"`
	StartTime   time.Time      `json:"start_time"`
	EndTime     time.Time      `json:"end_time"`
	Price       float64        `json:"price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
