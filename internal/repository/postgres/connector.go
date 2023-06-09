package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"siteAccess/internal/config"
	"siteAccess/internal/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db *sqlx.DB
}

func New(cfg config.Postgres) *Db {
	conn, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.User,
			cfg.Pass,
			cfg.Address,
			cfg.Port,
			cfg.Db,
		))
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: conn}
}

func (d *Db) GetTime(ctx context.Context, site string) (*domain.Answer, error) {
	var result domain.Answer
	err := d.db.GetContext(ctx, &result.Time, "SELECT response_time FROM access WHERE site = $1", site)
	if err == sql.ErrNoRows {
		result.Message = "n/a"
		return &result, nil
	}
	if err != nil {
		err = fmt.Errorf("get time failed, error: %w", err)
		return nil, err
	}

	return &result, nil
}

func (d *Db) GetMinTime(ctx context.Context) (*domain.Site, error) {
	var result domain.Site
	query := `SELECT site FROM access WHERE response_time = (SELECT MIN(response_time) FROM access)`
	err := d.db.GetContext(ctx, &result, query)
	if err != nil {
		err = fmt.Errorf("get site failed, error: %w", err)
		return nil, err
	}

	return &result, nil
}

func (d *Db) GetMaxTime(ctx context.Context) (*domain.Site, error) {
	var result domain.Site
	query := `SELECT site FROM access WHERE response_time = (SELECT MAX(response_time) FROM access)`
	err := d.db.GetContext(ctx, &result, query)
	if err != nil {
		err = fmt.Errorf("get site failed, error: %w", err)
		return nil, err
	}

	return &result, nil
}

func (d *Db) Update(ctx context.Context, newData map[string]time.Duration) error {
	query := `INSERT INTO access (site, response_time) VALUES ($1, $2) ON CONFLICT (site) DO UPDATE SET response_time = excluded.response_time`
	for site, duration := range newData {
		if duration == 0 {
			continue
		}
		_, err := d.db.ExecContext(ctx, query, site, duration.Milliseconds())
		if err != nil {
			err = fmt.Errorf("update data failed, error: %w", err)
			return err
		}
	}

	return nil
}
