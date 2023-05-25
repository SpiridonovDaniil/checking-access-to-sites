package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	router "siteAccess/internal/app/api/http"
	"siteAccess/internal/app/api/service"
	"siteAccess/internal/app/worker"
	"siteAccess/internal/config"
	"siteAccess/internal/repository/postgres"
	"time"
)

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

	service := service.New(db)
	r := router.NewServer(service)
	err := r.Listen(":" + cfg.Service.Port)
	if err != nil {
		panic(err)
	}

	http.Handle("/metrics", promhttp.Handler())
	rec := promauto.NewCounter(prometheus.CounterOpts{
		Name: "endpoint_registered",
		Help: "number of endpoints used",
		ConstLabels: map[string]string{
			"endpoint1": "site",
			"endpoint2": "min",
			"endpoint3": "max",
		},
	})
	go func() {

		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			panic(err)
		}
		rec.Inc()
	}()
}
