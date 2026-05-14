package repository

import (
	"Kevinmajesta/backend_bioskopMKP/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(schedule *entity.Schedule) (*entity.Schedule, error)
	UpdateSchedule(id_schedule uuid.UUID, schedule *entity.Schedule) (*entity.Schedule, error)
	DeleteSchedule(id_schedule uuid.UUID) error
	FindScheduleByID(id_schedule uuid.UUID) (*entity.Schedule, error)
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

func (r *scheduleRepository) UpdateSchedule(id_schedule uuid.UUID, schedule *entity.Schedule) (*entity.Schedule, error) {
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

	if schedule.IsCancelled != nil {
		fields["is_cancelled"] = *schedule.IsCancelled
	}

	if err := r.db.Model(&entity.Schedule{}).Where("id_schedule = ?", id_schedule).Updates(fields).Error; err != nil {
		return nil, err
	}
	return r.FindScheduleByID(id_schedule)
}

func (r *scheduleRepository) DeleteSchedule(id_schedule uuid.UUID) error {
	return r.db.Delete(&entity.Schedule{}, id_schedule).Error
}

func (r *scheduleRepository) FindScheduleByID(id_schedule uuid.UUID) (*entity.Schedule, error) {
	schedule := new(entity.Schedule)
	if err := r.db.Where("id_schedule = ?", id_schedule).First(schedule).Error; err != nil {
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
