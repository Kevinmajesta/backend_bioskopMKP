package main

import (
	"Kevinmajesta/backend_bioskopMKP/configs"
	"Kevinmajesta/backend_bioskopMKP/internal/builder"
	"Kevinmajesta/backend_bioskopMKP/pkg/postgres"
	"Kevinmajesta/backend_bioskopMKP/pkg/server"
	"Kevinmajesta/backend_bioskopMKP/pkg/token"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	tokenUseCase := token.NewTokenUseCase(cfg.JWTSecret)

	publicRoutes := builder.BuildPublicRoutes(db, tokenUseCase, cfg)
	privateRoutes := builder.BuildPrivateRoutes(db, cfg, tokenUseCase)

	srv := server.NewServer("app", publicRoutes, privateRoutes, cfg.JWTSecret)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
