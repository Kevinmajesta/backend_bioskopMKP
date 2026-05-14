package binder

import "time"

type ScheduleCreateRequest struct {
	MovieTitle string    `json:"movie_title" validate:"required"`
	CinemaName string    `json:"cinema_name" validate:"required"`
	StudioName string    `json:"studio_name" validate:"required"`
	StartTime  time.Time `json:"start_time" validate:"required"`
	EndTime    time.Time `json:"end_time" validate:"required"`
	Price      float64   `json:"price" validate:"required"`
}

type ScheduleUpdateRequest struct {
	MovieTitle  string    `json:"movie_title"`
	CinemaName  string    `json:"cinema_name"`
	StudioName  string    `json:"studio_name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Price       float64   `json:"price"`
	IsCancelled *bool     `json:"is_cancelled"` // Use pointer to allow false value in partial update
}
