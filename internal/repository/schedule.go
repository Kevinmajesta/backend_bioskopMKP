package repository

import (
	"Kevinmajesta/backend_bioskopMKP/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(schedule *entity.Schedule) (*entity.Schedule, error)
	UpdateSchedule(id uuid.UUID, schedule *entity.Schedule) (*entity.Schedule, error)
	DeleteSchedule(id uuid.UUID) error
	FindScheduleByID(id uuid.UUID) (*entity.Schedule, error)
	FindAllSchedules() ([]entity.Schedule, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db}
}

func (r *scheduleRepository) CreateSchedule(schedule *entity.Schedule) (*entity.Schedule, error) {
	if err := r.db.Create(schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *scheduleRepository) UpdateSchedule(id uuid.UUID, schedule *entity.Schedule) (*entity.Schedule, error) {
	// Menggunakan map agar GORM tidak mengabaikan zero-value (seperti false atau 0)
	fields := make(map[string]interface{})

	if schedule.MovieTitle != "" {
		fields["movie_title"] = schedule.MovieTitle
	}
	if schedule.CinemaName != "" {
		fields["cinema_name"] = schedule.CinemaName
	}
	if schedule.StudioName != "" {
		fields["studio_name"] = schedule.StudioName
	}
	if !schedule.StartTime.IsZero() {
		fields["start_time"] = schedule.StartTime
	}
	if !schedule.EndTime.IsZero() {
		fields["end_time"] = schedule.EndTime
	}
	if schedule.Price != 0 {
		fields["price"] = schedule.Price
	}

	// Untuk boolean (pointer), kita cek apakah nil atau tidak
	if schedule.IsCancelled != nil {
		fields["is_cancelled"] = *schedule.IsCancelled
	}

	if err := r.db.Model(&entity.Schedule{}).Where("id = ?", id).Updates(fields).Error; err != nil {
		return nil, err
	}
	return r.FindScheduleByID(id)
}

func (r *scheduleRepository) DeleteSchedule(id uuid.UUID) error {
	return r.db.Delete(&entity.Schedule{}, id).Error
}

func (r *scheduleRepository) FindScheduleByID(id uuid.UUID) (*entity.Schedule, error) {
	schedule := new(entity.Schedule)
	if err := r.db.Where("id = ?", id).First(schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *scheduleRepository) FindAllSchedules() ([]entity.Schedule, error) {
	var schedules []entity.Schedule
	if err := r.db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
