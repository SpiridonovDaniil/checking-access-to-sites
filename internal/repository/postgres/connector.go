package postgres

import (
	"context"
	"fmt"
	"log"
	"siteAccess/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"siteAccess/internal/config"
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

func (d *Db) GetTime(ctx context.Context, site string) (*domain.Time, error) {
	var result *domain.Time //TODO нужно ли делать транзакцию при зависимости использования методов и получения от них статистики?
	err := d.db.SelectContext(ctx, &result, "SELECT response_time FROM access WHERE site = $1", site)
	if err != nil {
		err = fmt.Errorf("get time failed, error: %w", err)
		return nil, err
	}

	_, err = d.db.ExecContext(ctx, "UPDATE statistic SET number_of_requests = number_of_requests + 1 WHERE endpoint = 'getTime'")
	if err != nil {
		err = fmt.Errorf("update statistics data failed, error: %w", err)
		return nil, err
	}

	return result, nil
}

func (d *Db) GetMinTime(ctx context.Context) (*domain.Site, error) {
	var result *domain.Site
	err := d.db.SelectContext(ctx, &result, "SELECT site FROM access WHERE MIN(response_time)")
	if err != nil {
		err = fmt.Errorf("get site failed, error: %w", err)
		return nil, err
	}

	_, err = d.db.ExecContext(ctx, "UPDATE statistic SET number_of_requests = number_of_requests + 1 WHERE endpoint = 'getMinTime'")
	if err != nil {
		err = fmt.Errorf("update statistics data failed, error: %w", err)
		return nil, err
	}
	return result, nil
}

func (d *Db) GetMaxTime(ctx context.Context) (*domain.Site, error) {
	var result *domain.Site
	err := d.db.SelectContext(ctx, &result, "SELECT site FROM access WHERE MAX(response_time)")
	if err != nil {
		err = fmt.Errorf("get site failed, error: %w", err)
		return nil, err
	}

	_, err = d.db.ExecContext(ctx, "UPDATE statistic SET number_of_requests = number_of_requests + 1 WHERE endpoint = 'getMaxTime'")
	if err != nil {
		err = fmt.Errorf("update statistics data failed, error: %w", err)
		return nil, err
	}
	return result, nil
}

func (d *Db) Update(ctx context.Context, newData map[string]time.Duration) error {
	return nil
}
