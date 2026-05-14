package repository

import (
	"Kevinmajesta/backend_bioskopMKP/internal/entity"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(schedule *entity.Schedule) (*entity.Schedule, error)
	UpdateSchedule(id uint, schedule *entity.Schedule) (*entity.Schedule, error)
	DeleteSchedule(id uint) error
	FindScheduleByID(id uint) (*entity.Schedule, error)
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

func (r *scheduleRepository) UpdateSchedule(id uint, schedule *entity.Schedule) (*entity.Schedule, error) {
	if err := r.db.Model(&entity.Schedule{}).Where("id = ?", id).Updates(schedule).Error; err != nil {
		return nil, err
	}
	return r.FindScheduleByID(id)
}

func (r *scheduleRepository) DeleteSchedule(id uint) error {
	return r.db.Delete(&entity.Schedule{}, id).Error
}

func (r *scheduleRepository) FindScheduleByID(id uint) (*entity.Schedule, error) {
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
