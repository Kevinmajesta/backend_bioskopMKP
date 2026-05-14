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

func AppPrivateRoutes() []*route.Route {
	return nil
}