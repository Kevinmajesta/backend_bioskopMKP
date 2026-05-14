package router

import (
	"net/http"

	"Kevinmajesta/backend_bioskopMKP/internal/http/handler"
	"Kevinmajesta/backend_bioskopMKP/pkg/route"
)

func AppPublicRoutes(userHandler *handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: userHandler.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
	}
}

func AppPrivateRoutes(scheduleHandler *handler.ScheduleHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/schedules",
			Handler: scheduleHandler.Create,
		},
		{
			Method:  http.MethodPut,
			Path:    "/schedules/:id_schedule",
			Handler: scheduleHandler.Update,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/schedules/:id_schedule",
			Handler: scheduleHandler.Delete,
		},
		{
			Method:  http.MethodGet,
			Path:    "/schedules/:id_schedule",
			Handler: scheduleHandler.GetByID,
		},
		{
			Method:  http.MethodGet,
			Path:    "/schedules",
			Handler: scheduleHandler.GetAll,
		},
	}
}
