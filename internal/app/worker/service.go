package worker

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"siteAccess/internal/config"
	"siteAccess/internal/repository/postgres"
	"sync"
	"time"
)

func Worker(cfg config.Site, db *postgres.Db) error {
	table := make(map[string]time.Duration)
	var wg sync.WaitGroup
	for _, site := range cfg.Site {
		site := site
		wg.Add(1)
		go func(url string) {
			now := time.Now()
			_, err := http.Get("https://" + site)
			difference := time.Since(now)
			if err != nil {
				difference = 0
			}
			table[site] = difference
			wg.Done()
		}(site)
	}
	wg.Wait()

	ctx := fiber.Ctx{} // где взять?
	err := db.Update(ctx.Context(), table)
	if err != nil {
		return err
	}

	return nil
}
