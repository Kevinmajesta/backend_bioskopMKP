package builder

import (
	"Kevinmajesta/backend_bioskopMKP/configs"
	"Kevinmajesta/backend_bioskopMKP/internal/http/handler"
	"Kevinmajesta/backend_bioskopMKP/internal/http/router"
	"Kevinmajesta/backend_bioskopMKP/internal/repository"
	"Kevinmajesta/backend_bioskopMKP/internal/service"
	"Kevinmajesta/backend_bioskopMKP/pkg/route"
	"Kevinmajesta/backend_bioskopMKP/pkg/token"

	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase,
	cfg *configs.Config) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, tokenUseCase)

	userHandler := handler.NewUserHandler(userService)

	return router.AppPublicRoutes(userHandler)
}

func BuildPrivateRoutes(db *gorm.DB, cfg *configs.Config, tokenUseCase token.TokenUseCase) []*route.Route {
	return router.AppPrivateRoutes()
}
