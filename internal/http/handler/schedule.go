package handler

import (
	"net/http"

	"Kevinmajesta/backend_bioskopMKP/internal/entity"
	"Kevinmajesta/backend_bioskopMKP/internal/http/binder"
	"Kevinmajesta/backend_bioskopMKP/internal/service"
	"Kevinmajesta/backend_bioskopMKP/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	scheduleService service.ScheduleService
}

func NewScheduleHandler(scheduleService service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{scheduleService}
}

func (h *ScheduleHandler) Create(c echo.Context) error {
	input := new(binder.ScheduleCreateRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid input"))
	}

	schedule := &entity.Schedule{
		MovieTitle: input.MovieTitle,
		CinemaName: input.CinemaName,
		StudioName: input.StudioName,
		StartTime:  input.StartTime,
		EndTime:    input.EndTime,
		Price:      input.Price,
	}

	result, err := h.scheduleService.CreateSchedule(schedule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse(http.StatusCreated, "schedule created", result))
}

func (h *ScheduleHandler) Update(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id_schedule"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid id format"))
	}

	input := new(binder.ScheduleUpdateRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid input"))
	}

	schedule := &entity.Schedule{
		MovieTitle: input.MovieTitle,
		CinemaName: input.CinemaName,
		StudioName: input.StudioName,
		StartTime:  input.StartTime,
		EndTime:    input.EndTime,
		Price:      input.Price,
	}

	if input.IsCancelled != nil {
		schedule.IsCancelled = input.IsCancelled
	}

	result, err := h.scheduleService.UpdateSchedule(id, schedule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "schedule updated", result))
}

func (h *ScheduleHandler) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id_schedule"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid id format"))
	}

	if err := h.scheduleService.DeleteSchedule(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "schedule deleted", nil))
}

func (h *ScheduleHandler) GetByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id_schedule"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid id format"))
	}

	result, err := h.scheduleService.GetScheduleByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "schedule not found"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success", result))
}

func (h *ScheduleHandler) GetAll(c echo.Context) error {
	results, err := h.scheduleService.GetAllSchedules()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success", results))
}
