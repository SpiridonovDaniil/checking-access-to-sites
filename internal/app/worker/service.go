package worker

import (
	"context"
	"net/http"
	"siteAccess/internal/config"
	"siteAccess/internal/repository/postgres"
	"sync"
	"time"
)

func Worker(ctx context.Context, cfgS config.Site, cfgU config.Interval, db *postgres.Db) error {
	table := make(map[string]time.Duration)
	var wg sync.WaitGroup
	for _, site := range cfgS.Site {
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

	ctx, cancel := context.WithTimeout(ctx, cfgU.DbTimeout)
	defer cancel()

	err := db.Update(ctx, table) //TODO передать указатель на карту?
	if err != nil {
		return err
	}

	return nil
}
