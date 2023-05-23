package main

import (
	router "siteAccess/internal/app/api/http"
	"siteAccess/internal/app/api/service"
	"siteAccess/internal/config"
	"siteAccess/internal/repository/postgres"
)

func main() {
	cfg := config.Read()

	db := postgres.New(cfg.Postgres)
	service := service.New(db)
	r := router.NewServer(service)
	err := r.Listen(":" + cfg.Service.Port)
	if err != nil {
		panic(err)
	}
}
