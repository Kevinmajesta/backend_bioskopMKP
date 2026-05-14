package service

import (
	"Kevinmajesta/backend_bioskopMKP/internal/entity"
	"Kevinmajesta/backend_bioskopMKP/internal/repository"
	"github.com/google/uuid"
)

type ScheduleService interface {
	CreateSchedule(schedule *entity.Schedule) (*entity.Schedule, error)
	UpdateSchedule(id uuid.UUID, schedule *entity.Schedule) (*entity.Schedule, error)
	DeleteSchedule(id uuid.UUID) error
	GetScheduleByID(id uuid.UUID) (*entity.Schedule, error)
	GetAllSchedules() ([]entity.Schedule, error)
}

type scheduleService struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleService(scheduleRepository repository.ScheduleRepository) ScheduleService {
	return &scheduleService{scheduleRepository}
}

func (s *scheduleService) CreateSchedule(schedule *entity.Schedule) (*entity.Schedule, error) {
	return s.scheduleRepository.CreateSchedule(schedule)
}

func (s *scheduleService) UpdateSchedule(id uuid.UUID, schedule *entity.Schedule) (*entity.Schedule, error) {
	return s.scheduleRepository.UpdateSchedule(id, schedule)
}

func (s *scheduleService) DeleteSchedule(id uuid.UUID) error {
	return s.scheduleRepository.DeleteSchedule(id)
}

func (s *scheduleService) GetScheduleByID(id uuid.UUID) (*entity.Schedule, error) {
	return s.scheduleRepository.FindScheduleByID(id)
}

func (s *scheduleService) GetAllSchedules() ([]entity.Schedule, error) {
	return s.scheduleRepository.FindAllSchedules()
}
