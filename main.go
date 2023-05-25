package main

import (
	"context"
	"log"
	"net/http"
	"time"

	_ "siteAccess/docs"
	router "siteAccess/internal/app/api/http"
	"siteAccess/internal/app/api/service"
	"siteAccess/internal/app/worker"
	"siteAccess/internal/config"
	"siteAccess/internal/repository/postgres"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title SiteAccess
// @version 1.0
// @description Swagger API for Golang Project siteAccess
// @termsOfService http://swagger.io/terms/
// @contact.name Daniil56
// @contact.email daniil13.spiridonov@yandex.ru
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.Read()

	ctx := context.Background()

	db := postgres.New(cfg.Postgres)

	go func() {
		var t time.Duration
		for {
			select {
			case <-time.After(t):
				t = cfg.Interval.ReloadWorker
				err := worker.Worker(ctx, cfg.Site, cfg.Interval, db)
				if err != nil {
					log.Println("[worker]", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	ro := http.NewServeMux()
	ro.Handle("/metrics", promhttp.Handler())
	rec := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "endpoint_registered",
		Help: "number of endpoints used",
	},
		[]string{"endpoints"},
	)
	go func() {
		err := http.ListenAndServe(":9090", ro)
		if err != nil {
			panic(err)
		}
	}()

	service := service.New(db)
	r := router.NewServer(service, rec)
	err := r.Listen(":" + cfg.Service.Port)
	if err != nil {
		panic(err)
	}
}
